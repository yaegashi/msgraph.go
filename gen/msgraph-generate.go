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
	"unicode"

	_ "github.com/rickb777/date/period"
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
	"Edm.Duration":         "Duration",
	"Edm.TimeOfDay":        "TimeOfDay",
	"Edm.Date":             "Date",
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
	Sym         string
	Members     []*EnumTypeMember
	Description string
}

type EnumTypeMember struct {
	Name        string
	Sym         string
	Value       string
	Description string
}

type EntityType struct {
	Name        string
	Sym         string
	Type        string
	Base        string
	Members     []*EntityTypeMember
	Navigations []*EntityTypeNavigation
	Description string
}

type EntityTypeMember struct {
	Name        string
	Sym         string
	Type        string
	Description string
}

type EntityTypeNavigation struct {
	Name        string
	Sym         string
	Type        string
	Description string
}

type ActionType struct {
	Name                 string
	Sym                  string
	BindingParameterType string
	Parameters           []*ActionTypeParameter
	ReturnType           string
	Description          string
}

type ActionTypeParameter struct {
	Name        string
	Sym         string
	Type        string
	Description string
}

type EntitySet struct {
	Name string
	Sym  string
	Type string
}

type Singleton struct {
	Name string
	Sym  string
	Type string
}

// lintName returns a different name if it should be different.
// Taken from https://github.com/golang/lint/blob/master/lint.go
func lintName(name string) (should string) {
	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word
		if i+1 == len(runes) {
			eow = true
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u, ok := commonInitialisms[strings.ToUpper(word)]; ok {
			// Keep consistent case, which is lowercase only at the start.
			if w == 0 && unicode.IsLower(runes[w]) {
				u = strings.ToLower(u)
			}
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))
		} else if w > 0 && strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
// Taken from https://github.com/golang/lint/blob/master/lint.go
var commonInitialisms = map[string]string{
	"ACL":    "ACL",
	"API":    "API",
	"ASCII":  "ASCII",
	"CPU":    "CPU",
	"CSS":    "CSS",
	"DNS":    "DNS",
	"EOF":    "EOF",
	"GUID":   "GUID",
	"HTML":   "HTML",
	"HTTP":   "HTTP",
	"HTTPS":  "HTTPS",
	"ICMP":   "ICMP",
	"ID":     "ID",
	"IDS":    "IDs",
	"IGMP":   "IGMP",
	"IOS":    "IOS",
	"IP":     "IP",
	"IPV4":   "IPv4",
	"IPV6":   "IPv6",
	"JSON":   "JSON",
	"LHS":    "LHS",
	"MFA":    "MFA",
	"OAUTH2": "OAuth2",
	"OMA":    "OMA",
	"QPS":    "QPS",
	"RAM":    "RAM",
	"RHS":    "RHS",
	"RPC":    "RPC",
	"SKU":    "SKU",
	"SKUS":   "SKUs",
	"SLA":    "SLA",
	"SMTP":   "SMTP",
	"SQL":    "SQL",
	"SSH":    "SSH",
	"TCP":    "TCP",
	"TLS":    "TLS",
	"TTL":    "TTL",
	"UDP":    "UDP",
	"UI":     "UI",
	"UID":    "UID",
	"URI":    "URI",
	"URL":    "URL",
	"UTF8":   "UTF8",
	"UUID":   "UUID",
	"VM":     "VM",
	"XML":    "XML",
	"XMPP":   "XMPP",
	"XSRF":   "XSRF",
	"XSS":    "XSS",
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

func exported(n string) string {
	return lintName(strings.Title(n))
}

func isCollectionType(t string) bool {
	return strings.HasPrefix(t, colPrefix)
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
	Created               map[string]bool
	SymTypeMap            map[string]string
	X, Y, Z               interface{}
}

func (g *Generator) SymBaseType(t string) string {
	if x, ok := g.SymTypeMap[t]; ok {
		return x
	}
	if x, ok := stripNSPrefix(t); ok {
		return exported(x)
	}
	if strings.HasPrefix(t, colPrefix) {
		return g.SymBaseType(t[len(colPrefix) : len(t)-1])
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func (g *Generator) SymFromType(t string) string {
	if x, ok := g.SymTypeMap[t]; ok {
		return x
	}
	if x, ok := stripNSPrefix(t); ok {
		return exported(x)
	}
	if strings.HasPrefix(t, colPrefix) {
		return g.SymBaseType(t[len(colPrefix):len(t)-1]) + "Collection"
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func (g *Generator) TypeFromType(t string) string {
	if x, ok := reservedTypeTable[t]; ok {
		return ptrType(x)
	}
	if x, ok := g.SymTypeMap[t]; ok {
		return ptrType(x)
	}
	if x, ok := stripNSPrefix(t); ok {
		return ptrType(exported(x))
	}
	if strings.HasPrefix(t, colPrefix) {
		return "[]" + g.TypeFromType(t[len(colPrefix) : len(t)-1])[1:]
	}
	panic(fmt.Errorf("Unknown type %s", t))
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
	actionRequestBuilderMap := map[string][]string{}

	for _, x := range schema.Elems {
		switch x.XMLName.Local {
		case "EnumType":
			m := attrMap(x.Attrs)
			n := m["Name"]
			t := &EnumType{Name: n, Sym: exported(n), Description: "undocumented"}
			for _, y := range x.Elems {
				n := y.Attrs[0].Value
				// ex. Collection(String) -> StringCollection
				if strings.HasPrefix(n, colPrefix) {
					n = n[len(colPrefix):len(n)-1] + "Collection"
				}
				v := y.Attrs[1].Value
				m := &EnumTypeMember{Name: n, Sym: exported(n), Value: v, Description: "undocumented"}
				t.Members = append(t.Members, m)
			}
			enumTypeMap[n] = t
		case "EntityType", "ComplexType":
			m := attrMap(x.Attrs)
			n := m["Name"]
			if _, ok := reservedTypeTable[n]; ok {
				continue
			}
			t := nsPrefix + n
			b, _ := m["BaseType"]
			et := &EntityType{Name: n, Sym: exported(n), Type: t, Base: b, Description: "undocumented"}
			if strings.HasSuffix(et.Sym, "Request") {
				et.Sym += "Object"
				g.SymTypeMap[t] = et.Sym
			}
			for _, y := range x.Elems {
				ma := attrMap(y.Attrs)
				switch y.XMLName.Local {
				case "Property":
					n := ma["Name"]
					t := ma["Type"]
					m := &EntityTypeMember{Name: n, Sym: exported(n), Type: t, Description: "undocumented"}
					et.Members = append(et.Members, m)
				case "NavigationProperty":
					n := ma["Name"]
					t := ma["Type"]
					m := &EntityTypeNavigation{Name: n, Sym: exported(n), Type: t, Description: "undocumented"}
					if strings.HasSuffix(m.Sym, "Request") {
						m.Sym += "Navigation"
					}
					et.Navigations = append(et.Navigations, m)
				}
			}
			entityTypeMap[et.Name] = et
		case "Action":
			m := attrMap(x.Attrs)
			n := m["Name"]
			at := &ActionType{Name: n, Sym: exported(n), Description: "undocumented"}
			if strings.HasSuffix(at.Sym, "Request") {
				at.Sym += "Action"
			}
			for _, y := range x.Elems {
				ma := attrMap(y.Attrs)
				switch y.XMLName.Local {
				case "Parameter":
					n := ma["Name"]
					t := ma["Type"]
					m := &ActionTypeParameter{Name: n, Sym: exported(n), Type: t, Description: "undocumented"}
					at.Parameters = append(at.Parameters, m)
				case "ReturnType":
					at.ReturnType = ma["Type"]
				}
			}
			at.BindingParameterType = at.Parameters[0].Type
			at.Parameters = at.Parameters[1:]
			actionTypeMap[at.BindingParameterType] = append(actionTypeMap[at.BindingParameterType], at)
		case "EntityContainer":
			for _, y := range x.Elems {
				ma := attrMap(y.Attrs)
				switch y.XMLName.Local {
				case "EntitySet":
					s := &EntitySet{
						Name: ma["Name"],
						Sym:  exported(ma["Name"]),
						Type: ma["EntityType"],
					}
					entitySetMap[s.Name] = s
				case "Singleton":
					s := &Singleton{
						Name: ma["Name"],
						Sym:  exported(ma["Name"]),
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

	var out io.WriteCloser

	for _, path := range []string{"msgraph", "extensions"} {
		out, err = g.Create(path)
		if err != nil {
			return err
		}
		err = tmpl.ExecuteTemplate(out, path+".go.tmpl", g)
		if err != nil {
			return err
		}
		out.Close()
	}

	keys := []string{}
	for x, _ := range enumTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		x := enumTypeMap[key]
		out, err = g.Create(x.Sym)
		g.X = x
		if err != nil {
			return err
		}
		err := tmpl.ExecuteTemplate(out, "enum.go.tmpl", g)
		if err != nil {
			return err
		}
		out.Close()
	}

	keys = nil
	for x, _ := range entityTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		x := entityTypeMap[key]
		out, err = g.Create(x.Sym)
		if err != nil {
			return err
		}
		g.X = x
		err := tmpl.ExecuteTemplate(out, "model.go.tmpl", g)
		if err != nil {
			return err
		}
		out.Close()
	}

	keys = nil
	for x, _ := range actionTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, a := range keys {
		if _, ok := reservedTypeTable[a]; ok {
			continue
		}
		x := actionTypeMap[a]
		out, err = g.Create(g.SymBaseType(a))
		if err != nil {
			return err
		}
		requestModelMap[g.SymBaseType(a)] = true
		for _, y := range x {
			g.X = y
			err := tmpl.ExecuteTemplate(out, "action.go.tmpl", g)
			if err != nil {
				return err
			}
		}
		out.Close()
	}

	for _, x := range entityTypeMap {
		if len(x.Navigations) == 0 {
			continue
		}
		requestModelMap[x.Sym] = true
		for _, y := range x.Navigations {
			requestModelMap[g.SymBaseType(y.Type)] = true
		}
	}
	for _, x := range entitySetMap {
		requestModelMap[g.SymBaseType(x.Type)] = true
	}
	for _, x := range singletonMap {
		requestModelMap[g.SymBaseType(x.Type)] = true
	}

	keys = nil
	for x, _ := range requestModelMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, x := range keys {
		out, err = g.Create(x)
		if err != nil {
			return err
		}
		g.X = x
		err := tmpl.ExecuteTemplate(out, "request_model.go.tmpl", g)
		if err != nil {
			return err
		}
		out.Close()
	}

	out, err = g.Create("GraphService")
	if err != nil {
		return err
	}
	keys = nil
	for x, _ := range entitySetMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		g.X = &EntityType{Name: "GraphService", Sym: "GraphService"}
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
		g.X = &EntityType{Name: "GraphService", Sym: "GraphService"}
		g.Y = singletonMap[key]
		err := tmpl.ExecuteTemplate(out, "request_navigation.go.tmpl", g)
		if err != nil {
			return err
		}
	}
	out.Close()

	keys = nil
	for x, _ := range entityTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		x := entityTypeMap[key]
		actionRequestBuilderMap[x.Type] = append(actionRequestBuilderMap[x.Type], x.Sym)
		if len(x.Navigations) == 0 {
			continue
		}
		out, err = g.Create(x.Sym)
		if err != nil {
			return err
		}
		g.X = x
		sort.Slice(x.Navigations, func(i, j int) bool { return x.Navigations[i].Name < x.Navigations[j].Name })
		for _, y := range x.Navigations {
			g.Y = y
			if isCollectionType(y.Type) {
				actionRequestBuilderMap[y.Type] = append(actionRequestBuilderMap[y.Type], x.Sym+y.Sym+"Collection")
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
		out.Close()
	}

	keys = nil
	for x, _ := range actionTypeMap {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, a := range keys {
		x := actionTypeMap[a]
		if _, ok := reservedTypeTable[a]; ok {
			continue
		}
		out, err = g.Create(g.SymBaseType(a))
		if err != nil {
			return err
		}
		for _, y := range x {
			g.Y = y
			if b, ok := actionRequestBuilderMap[a]; ok {
				g.X = b
				if y.ReturnType == "" {
					err = tmpl.ExecuteTemplate(out, "request_action_void.go.tmpl", g)
				} else if isCollectionType(y.ReturnType) {
					err = tmpl.ExecuteTemplate(out, "request_action_collection.go.tmpl", g)
				} else {
					err = tmpl.ExecuteTemplate(out, "request_action_single.go.tmpl", g)
				}
				if err != nil {
					return err
				}
			}
		}
		out.Close()
	}

	return nil
}

func (g *Generator) Clean() error {
	log.Printf("Creating directory %s", g.Out)
	err := os.MkdirAll(filepath.Dir(g.Out), 0755)
	if err != nil {
		return err
	}
	dir, err := os.Open(g.Out)
	if err != nil {
		return err
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		if strings.HasSuffix(file.Name(), "_test.go") {
			continue
		}
		path := filepath.Join(g.Out, file.Name())
		log.Printf("Removing %s", path)
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) Create(path string) (io.WriteCloser, error) {
	if !strings.HasSuffix(path, ".go") {
		path += ".go"
	}
	path = filepath.Join(g.Out, path)
	if g.Created[path] {
		return os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	}
	log.Printf("Creating %s", path)
	g.Created[path] = true
	out, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	out.WriteString("// Code generated by msgraph-generate.go DO NOT EDIT.\n\npackage msgraph\n\n")
	return out, nil
}

func (g *Generator) Format() error {
	var files []string
	for file := range g.Created {
		files = append(files, file)
	}
	log.Printf("Formatting %s", strings.Join(files, " "))
	args := append([]string{"-w"}, files...)
	return exec.Command(g.Fmt, args...).Run()
}

func main() {
	g := &Generator{Created: map[string]bool{}, SymTypeMap: map[string]string{}}

	flag.StringVar(&g.BaseURL, "baseURL", "https://graph.microsoft.com/v1.0", "Base URL")
	flag.StringVar(&g.In, "in", "metadata-v1.0.xml", "Input file name")
	flag.StringVar(&g.Out, "out", "out", "Output folder name")
	flag.StringVar(&g.Fmt, "fmt", "goimports", "Formatter")
	flag.Parse()

	err := g.Clean()
	if err != nil {
		log.Fatalf("Failed to clean: %s", err)
	}

	err = g.Generate()
	if err != nil {
		log.Fatalf("Failed to generate: %s", err)
	}

	err = g.Format()
	if err != nil {
		log.Fatalf("Failed to format: %s", err)
	}
}
