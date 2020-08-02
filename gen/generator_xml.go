package main

import "encoding/xml"

type Elem struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Elems   []Elem     `xml:",any"`
}

func (e *Elem) AttrMap() map[string]string {
	m := map[string]string{}
	for _, x := range e.Attrs {
		m[x.Name.Local] = x.Value
	}
	return m
}
