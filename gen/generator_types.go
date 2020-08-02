package main

import (
	"fmt"
	"strings"
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
