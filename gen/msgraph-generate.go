// +build ignore

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"
)

const (
	nsPrefix  = "microsoft.graph."
	colPrefix = "Collection("
)

var reservedTypeTable = map[string]string{
	"Edm.Boolean":          "bool",
	"Edm.Byte":             "byte",
	"Edm.Int16":            "int",
	"Edm.Int32":            "int",
	"Edm.Int64":            "int",
	"Edm.Decimal":          "int",
	"Edm.Single":           "float64",
	"Edm.Double":           "float64",
	"Edm.Binary":           "Binary",
	"Edm.Stream":           "Stream",
	"Edm.Guid":             "UUID",
	"Edm.String":           "string",
	"Edm.DateTimeOffset":   "time.Time",
	"Edm.Duration":         "time.Duration",
	"Edm.TimeOfDay":        "time.Time",
	"Edm.Date":             "time.Time",
	"microsoft.graph.Json": "json.RawMessage",
}

type Const struct {
	Name, Value, Type string
}

type Elem struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Elems   []Elem     `xml:",any"`
}

type EnumType struct {
	Name        string
	Members     []*EnumTypeMember
	Description string
}

type EnumTypeMember struct {
	Name        string
	Value       string
	Description string
}

type EntityType struct {
	Name        string
	Base        string
	Members     []*EntityTypeMember
	Navigations []*EntityTypeNavigation
	Description string
}

type EntityTypeMember struct {
	Name        string
	Type        string
	Description string
}

type EntityTypeNavigation struct {
	Name        string
	Type        string
	Description string
}

type EntitySet struct {
	Name string
	Type string
}

type Singleton struct {
	Name string
	Type string
}

func stripNSPrefix(t string) (string, bool) {
	ok := strings.HasPrefix(t, nsPrefix)
	if ok {
		t = t[len(nsPrefix):]
	}
	return t, ok
}

func symExported(n string) string {
	return strings.Title(n)
}

func typeConverted(t string) string {
	if val, ok := reservedTypeTable[t]; ok {
		return val
	}
	if t, ok := stripNSPrefix(t); ok {
		return symExported(t)
	}
	if strings.HasPrefix(t, colPrefix) {
		return "[]" + typeConverted(t[len(colPrefix):len(t)-1])
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func ptrTypeConverted(t string) string {
	t = typeConverted(t)
	switch t {
	case "json.RawMessage":
		return t
	}
	if strings.HasPrefix(t, "[]") {
		return t
	}
	return "*" + t
}

func symTypeConverted(t string) string {
	t = typeConverted(t)
	if strings.HasPrefix(t, "[]") {
		return "Collection" + t[2:]
	}
	return t
}

func attrMap(a []xml.Attr) map[string]string {
	m := map[string]string{}
	for _, x := range a {
		m[x.Name.Local] = x.Value
	}
	return m
}

var tmplX = struct {
	SymExported, TypeConverted, PtrTypeConverted, SymTypeConverted func(string) string
	Environ                                                        map[string]string
	Data, X                                                        interface{}
}{
	SymExported:      symExported,
	TypeConverted:    typeConverted,
	PtrTypeConverted: ptrTypeConverted,
	SymTypeConverted: symTypeConverted,
	Environ:          map[string]string{},
}

func main() {
	data := struct{ BaseURL, In, Out, Fmt string }{}
	tmplX.Data = &data

	for _, kv := range os.Environ() {
		s := strings.Split(kv, "=")
		tmplX.Environ[s[0]] = s[1]
	}

	flag.StringVar(&data.BaseURL, "baseURL", "https://graph.microsoft.com/v1.0", "Base URL")
	flag.StringVar(&data.In, "in", "metadata-v1.0.xml", "Input file name")
	flag.StringVar(&data.Out, "out", "-", "Output file name")
	flag.StringVar(&data.Fmt, "fmt", "goimports", "Formatter")
	flag.Parse()

	tmpl, err := template.ParseGlob("*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	inFile, err := os.Open(data.In)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	var o Elem
	dec := xml.NewDecoder(inFile)
	err = dec.Decode(&o)
	if err != nil {
		log.Fatal(err)
	}

	schema := o.Elems[0].Elems[0]
	enumTypeMap := map[string]*EnumType{}
	entityTypeMap := map[string]*EntityType{}
	entitySetMap := map[string]*EntitySet{}
	singletonMap := map[string]*Singleton{}
	baseTypeMap := map[string]bool{}

	for _, x := range schema.Elems {
		switch x.XMLName.Local {
		case "EnumType":
			m := attrMap(x.Attrs)
			n := m["Name"]
			t := &EnumType{
				Name:        n,
				Description: "undocumented",
			}
			for _, y := range x.Elems {
				n := symExported(y.Attrs[0].Value)
				v := y.Attrs[1].Value
				t.Members = append(t.Members, &EnumTypeMember{Name: n, Value: v, Description: "undocumented"})
			}
			enumTypeMap[n] = t
		case "EntityType", "ComplexType":
			m := attrMap(x.Attrs)
			n := m["Name"]
			if _, ok := reservedTypeTable[n]; ok {
				continue
			}
			b, ok := m["BaseType"]
			if ok {
				if bt, ok := stripNSPrefix(b); ok {
					baseTypeMap[bt] = true
				}
			}
			et := &EntityType{
				Name:        n,
				Base:        b,
				Description: "undocumented",
			}
			for _, y := range x.Elems {
				ma := attrMap(y.Attrs)
				switch y.XMLName.Local {
				case "Property":
					n := ma["Name"]
					t := ma["Type"]
					et.Members = append(et.Members, &EntityTypeMember{Name: n, Type: t, Description: "undocumented"})
				case "NavigationProperty":
					n := ma["Name"]
					t := ma["Type"]
					et.Navigations = append(et.Navigations, &EntityTypeNavigation{Name: n, Type: t, Description: "undocumented"})
				}
			}
			entityTypeMap[et.Name] = et
		case "EntityContainer":
			for _, y := range x.Elems {
				ma := attrMap(y.Attrs)
				switch y.XMLName.Local {
				case "EntitySet":
					s := &EntitySet{
						Name: ma["Name"],
						Type: ma["EntityType"],
					}
					entitySetMap[s.Name] = s
				case "Singleton":
					s := &Singleton{
						Name: ma["Name"],
						Type: ma["Type"],
					}
					singletonMap[s.Name] = s
				}
			}
		case "Annotations":
			mas := attrMap(x.Attrs)
			target, _ := stripNSPrefix(mas["Target"])
			targetList := strings.Split(target, "/")
			targetMember := ""
			if len(targetList) > 1 {
				target = targetList[0]
				targetMember = targetList[1]
			}
			for _, y := range x.Elems {
				switch y.XMLName.Local {
				case "Annotation":
					ma := attrMap(y.Attrs)
					term, _ := ma["Term"]
					str, _ := ma["String"]
					if term == "Org.OData.Core.V1.Description" {
						if e, ok := entityTypeMap[target]; ok {
							if targetMember == "" {
								e.Description = str
							} else {
								for _, mem := range e.Members {
									if mem.Name == targetMember {
										mem.Description = str
									}
								}
							}
						}
					}
				}
			}
		}
	}

	cmd := exec.Command(data.Fmt)
	out, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Stderr = os.Stderr
	if data.Out == "-" {
		cmd.Stdout = os.Stdout
	} else {
		cmd.Stdout, err = os.Create(data.Out)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Generating %s from %s", data.Out, data.In)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(out, "msgraph.go.tmpl", tmplX)
	if err != nil {
		log.Fatal(err)
	}

	keys := []string{}
	for x, _ := range enumTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		tmplX.X = enumTypeMap[key]
		err := tmpl.ExecuteTemplate(out, "enum.go.tmpl", tmplX)
		if err != nil {
			log.Fatal(err)
		}
	}

	keys = nil
	for x, _ := range entityTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		tmplX.X = entityTypeMap[key]
		err := tmpl.ExecuteTemplate(out, "entity.go.tmpl", tmplX)
		if err != nil {
			log.Fatal(err)
		}
	}

	serviceStructMap := map[string]bool{}
	for _, x := range entityTypeMap {
		if len(x.Navigations) == 0 {
			continue
		}
		xName := symExported(x.Name)
		serviceStructMap[xName] = true
		for _, y := range x.Navigations {
			yType := symTypeConverted(y.Type)
			serviceStructMap[yType] = true
		}
	}
	for _, x := range entitySetMap {
		xType := typeConverted(x.Type)
		serviceStructMap[xType] = true
		serviceStructMap["Collection"+xType] = true
	}
	for _, x := range singletonMap {
		xType := typeConverted(x.Type)
		serviceStructMap[xType] = true
	}

	keys = nil
	for x, _ := range serviceStructMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, x := range keys {
		tmplX.X = x
		err := tmpl.ExecuteTemplate(out, "service_struct.go.tmpl", tmplX)
		if err != nil {
			log.Fatal(err)
		}
	}

	keys = nil
	for x, _ := range entitySetMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		tmplX.X = entitySetMap[key]
		err := tmpl.ExecuteTemplate(out, "service_entity_set.go.tmpl", tmplX)
		if err != nil {
			log.Fatal(err)
		}
	}

	keys = nil
	for x, _ := range singletonMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		tmplX.X = singletonMap[key]
		err := tmpl.ExecuteTemplate(out, "service_singleton.go.tmpl", tmplX)
		if err != nil {
			log.Fatal(err)
		}
	}

	keys = nil
	for x, _ := range entityTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		x := entityTypeMap[key]
		tmplX.X = x
		xSymType := symExported(x.Name)
		sort.Slice(x.Navigations, func(i, j int) bool { return x.Navigations[i].Name < x.Navigations[j].Name })
		err := tmpl.ExecuteTemplate(out, "service_navigation.go.tmpl", tmplX)
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := serviceStructMap[xSymType]; ok {
			err = tmpl.ExecuteTemplate(out, "service_request.go.tmpl", tmplX)
			if err != nil {
				log.Fatal(err)
			}
		}
		if _, ok := serviceStructMap["Collection"+xSymType]; ok {
			err = tmpl.ExecuteTemplate(out, "service_collection_request.go.tmpl", tmplX)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	out.Close()

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
