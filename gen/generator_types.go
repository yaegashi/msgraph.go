package main

import (
	"fmt"
	"strings"
)

const (
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

func ptrType(t string) string {
	switch t {
	case "json.RawMessage":
		return t
	}
	return "*" + t
}

func sepNamespaceAndType(t string) (string, string) {
	parts := strings.Split(t, ".")
	if strings.HasPrefix(t, "Collection") {
		fmt.Println()
	}
	var ns string
	if len(parts) > 1 {
		ns = strings.Join(parts[:len(parts)-1], ".")
	}
	return ns, parts[len(parts)-1]
}

func exported(n string) string {
	return lintName(strings.Title(n))
}

func isCollectionType(t string) bool {
	return strings.HasPrefix(t, colPrefix)
}

func (g *Generator) SymBaseType(t string) string {
	if strings.HasPrefix(t, colPrefix) {
		return g.SymBaseType(t[len(colPrefix) : len(t)-1])
	}
	ns, typ := sepNamespaceAndType(t)
	if x, ok := g.SymTypeMap[ns][typ]; ok {
		return x
	}
	if ns != "" {
		return exported(typ)
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func (g *Generator) SymFromType(t string) string {
	if strings.HasPrefix(t, colPrefix) {
		return g.SymBaseType(t[len(colPrefix):len(t)-1]) + "Collection"
	}
	ns, typ := sepNamespaceAndType(t)
	tm := g.SymTypeMap[ns]
	if x, ok := tm[typ]; ok {
		return x
	}
	if ns != "" {
		return exported(typ)
	}
	panic(fmt.Errorf("Unknown type %s", t))
}

func (g *Generator) TypeFromType(t string) string {
	if x, ok := reservedTypeTable[t]; ok {
		return ptrType(x)
	}

	if strings.HasPrefix(t, colPrefix) {
		return "[]" + g.TypeFromType(t[len(colPrefix) : len(t)-1])[1:]
	}
	ns, typ := sepNamespaceAndType(t)
	tm := g.SymTypeMap[ns]
	if x, ok := tm[typ]; ok {
		return ptrType(x)
	}
	if ns != "" {
		return ptrType(exported(typ))
	}
	panic(fmt.Errorf("Unknown type %s", t))
}
