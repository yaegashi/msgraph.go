// +build ignore

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

type ActionType struct {
	Name                 string
	BindingParameterType string
	Parameters           []*ActionTypeParameter
	ReturnType           string
	Description          string
}

type ActionTypeParameter struct {
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

func ptrType(t string) string {
	switch t {
	case "json.RawMessage":
		return t
	}
	return "*" + t
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

func isCollectionType(t string) bool {
	return strings.HasPrefix(t, colPrefix)
}

func symBaseType(t string) string {
	if t[:2] == "[]" {
		t = t[2:]
	}
	if t[:1] == "*" {
		t = t[1:]
	}
	return t
}

func symCollectionType(t string) string {
	return "Collection" + symBaseType(t)
}

func attrMap(a []xml.Attr) map[string]string {
	m := map[string]string{}
	for _, x := range a {
		m[x.Name.Local] = x.Value
	}
	return m
}

type Generator struct {
	BaseURL, In, Out, Fmt string
	Environ               map[string]string
	Created               []string
	X, Y, Z               interface{}
}

func (g *Generator) TypeFromName(n string) string {
	return ptrType(symExported(n))
}

func (g *Generator) TypeFromType(t string) string {
	if val, ok := reservedTypeTable[t]; ok {
		return ptrType(val)
	}
	if t, ok := stripNSPrefix(t); ok {
		return ptrType(symExported(t))
	}
	if strings.HasPrefix(t, colPrefix) {
		return "[]" + g.TypeFromType(t[len(colPrefix) : len(t)-1])[1:]
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func (g *Generator) SymFromName(n string) string {
	return symExported(n)
}

func (g *Generator) SymFromType(t string) string {
	t = g.TypeFromType(t)
	if t[:2] == "[]" {
		return t[2:]
	}
	if t[:1] == "*" {
		return t[1:]
	}
	return t
}

func (g *Generator) Create(path string) (io.WriteCloser, error) {
	path = filepath.Join(g.Out, path)
	g.Created = append(g.Created, path)
	log.Printf("Creating %s", path)
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(path)
}

func (g *Generator) Generate() error {
	tmpl, err := template.ParseGlob("*.tmpl")
	if err != nil {
		return err
	}

	inFile, err := os.Open(g.In)
	if err != nil {
		return err
	}
	defer inFile.Close()

	var o Elem
	dec := xml.NewDecoder(inFile)
	err = dec.Decode(&o)
	if err != nil {
		return err
	}

	schema := o.Elems[0].Elems[0]
	enumTypeMap := map[string]*EnumType{}
	entityTypeMap := map[string]*EntityType{}
	actionTypeMap := map[string][]*ActionType{}
	entitySetMap := map[string]*EntitySet{}
	singletonMap := map[string]*Singleton{}
	requestModelMap := map[string]bool{}

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
			b, _ := m["BaseType"]
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
		case "Action":
			m := attrMap(x.Attrs)
			n := m["Name"]
			at := &ActionType{
				Name:        n,
				Description: "undocumented",
			}
			for _, y := range x.Elems {
				ma := attrMap(y.Attrs)
				switch y.XMLName.Local {
				case "Parameter":
					n := ma["Name"]
					t := ma["Type"]
					at.Parameters = append(at.Parameters, &ActionTypeParameter{Name: n, Type: t, Description: "undocumented"})
				case "ReturnType":
					at.ReturnType = ma["Type"]
				}
			}
			at.BindingParameterType = at.Parameters[0].Type
			at.Parameters = at.Parameters[1:]
			bType := g.TypeFromType(at.BindingParameterType)
			actionTypeMap[bType] = append(actionTypeMap[bType], at)
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

	out, err := g.Create("msgraph.go")
	if err != nil {
		return err
	}
	defer out.Close()

	err = tmpl.ExecuteTemplate(out, "msgraph.go.tmpl", g)
	if err != nil {
		return err
	}

	keys := []string{}
	for x, _ := range enumTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		g.X = enumTypeMap[key]
		err := tmpl.ExecuteTemplate(out, "enum.go.tmpl", g)
		if err != nil {
			return err
		}
	}

	keys = nil
	for x, _ := range entityTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		g.X = entityTypeMap[key]
		err := tmpl.ExecuteTemplate(out, "model.go.tmpl", g)
		if err != nil {
			return err
		}
	}

	for _, x := range actionTypeMap {
		requestModelMap[g.SymFromType(x[0].BindingParameterType)] = true
		for _, y := range x {
			g.X = y
			err := tmpl.ExecuteTemplate(out, "action.go.tmpl", g)
			if err != nil {
				return err
			}
		}
	}

	for _, x := range entityTypeMap {
		if len(x.Navigations) == 0 {
			continue
		}
		requestModelMap[g.SymFromName(x.Name)] = true
		for _, y := range x.Navigations {
			yType := g.TypeFromType(y.Type)
			requestModelMap[symBaseType(yType)] = true
		}
	}
	for _, x := range entitySetMap {
		requestModelMap[g.SymFromType(x.Type)] = true
	}
	for _, x := range singletonMap {
		requestModelMap[g.SymFromType(x.Type)] = true
	}

	keys = nil
	for x, _ := range requestModelMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, x := range keys {
		g.X = x
		err := tmpl.ExecuteTemplate(out, "request_model.go.tmpl", g)
		if err != nil {
			return err
		}
	}

	keys = nil
	for x, _ := range entitySetMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		g.X = &EntityType{Name: "GraphService"}
		g.Y = entitySetMap[key]
		err := tmpl.ExecuteTemplate(out, "request_collection_navigation.go.tmpl", g)
		if err != nil {
			return err
		}
	}

	keys = nil
	for x, _ := range singletonMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		g.X = &EntityType{Name: "GraphService"}
		g.Y = singletonMap[key]
		err := tmpl.ExecuteTemplate(out, "request_navigation.go.tmpl", g)
		if err != nil {
			return err
		}
	}

	actionRequestBuilderMap := map[string][]string{}

	keys = nil
	for x, _ := range entityTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		x := entityTypeMap[key]
		g.X = x
		sort.Slice(x.Navigations, func(i, j int) bool { return x.Navigations[i].Name < x.Navigations[j].Name })
		for _, y := range x.Navigations {
			g.Y = y
			yType := g.TypeFromType(y.Type)
			if isCollectionType(y.Type) {
				actionRequestBuilderMap[yType] = append(actionRequestBuilderMap[yType], g.SymFromName(x.Name)+g.SymFromName(y.Name)+"Collection")
				err := tmpl.ExecuteTemplate(out, "request_collection_navigation.go.tmpl", g)
				if err != nil {
					return err
				}
			} else {
				err := tmpl.ExecuteTemplate(out, "request_navigation.go.tmpl", g)
				if err != nil {
					return err
				}
			}
		}
		xType := g.TypeFromName(x.Name)
		actionRequestBuilderMap[xType] = append(actionRequestBuilderMap[xType], g.SymFromName(x.Name))
	}

	for x, y := range actionTypeMap {
		for _, z := range y {
			g.Y = z
			if a, ok := actionRequestBuilderMap[x]; ok {
				g.X = a
				if z.ReturnType == "" {
					err = tmpl.ExecuteTemplate(out, "request_action_void.go.tmpl", g)
				} else if isCollectionType(z.ReturnType) {
					err = tmpl.ExecuteTemplate(out, "request_action_collection.go.tmpl", g)
				} else {
					err = tmpl.ExecuteTemplate(out, "request_action_single.go.tmpl", g)
				}
				if err != nil {
					return err
				}
			}
		}
	}
	// keys = nil
	// for x, _ := range actionTypeMap {
	// 	keys = append(keys, x)
	// }
	// sort.Strings(keys)
	// for _, key := range keys {
	// 	x := actionTypeMap[key]
	// 	g.X = x
	// 	if x.ReturnType == "" {
	// 		err = tmpl.ExecuteTemplate(out, "action_void.go.tmpl", g)
	// 	} else {
	// 		aType := g.TypeFromType(x.ReturnType)
	// 		if strings.HasPrefix(aType, "[]") {
	// 			err = tmpl.ExecuteTemplate(out, "action_collection.go.tmpl", g)
	// 		} else {
	// 			err = tmpl.ExecuteTemplate(out, "action_single.go.tmpl", g)
	// 		}
	// 	}
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (g *Generator) Format() error {
	log.Printf("Formatting %s", strings.Join(g.Created, " "))
	args := append([]string{"-w"}, g.Created...)
	return exec.Command(g.Fmt, args...).Run()
}

func main() {
	g := &Generator{
		Environ: map[string]string{},
	}

	for _, kv := range os.Environ() {
		s := strings.Split(kv, "=")
		g.Environ[s[0]] = s[1]
	}

	flag.StringVar(&g.BaseURL, "baseURL", "https://graph.microsoft.com/v1.0", "Base URL")
	flag.StringVar(&g.In, "in", "metadata-v1.0.xml", "Input file name")
	flag.StringVar(&g.Out, "out", "out", "Output folder name")
	flag.StringVar(&g.Fmt, "fmt", "goimports", "Formatter")
	flag.Parse()

	err := g.Generate()
	if err != nil {
		log.Fatalf("Failed to generate: %s", err)
	}

	err = g.Format()
	if err != nil {
		log.Fatalf("Failed to format: %s", err)
	}
}
