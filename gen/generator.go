package main

type Generator struct {
	BaseURL, In, Out, Fmt string
	Created               map[string]bool
	SymTypeMap            map[string]string
	X, Y, Z               interface{}
}
