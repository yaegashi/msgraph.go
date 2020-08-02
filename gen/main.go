package main

//go:generate go run ./metadata -pretty -baseURL https://graph.microsoft.com/v1.0 -out metadata/v1.0.xml
//go:generate go run ./metadata -pretty -baseURL https://graph.microsoft.com/beta -out metadata/beta.xml
//go:generate go run . -baseURL https://graph.microsoft.com/v1.0 -in metadata/v1.0.xml -out ../v1.0
//go:generate go run . -baseURL https://graph.microsoft.com/beta -in metadata/beta.xml -out ../beta

import (
	"flag"
	"log"

	_ "github.com/rickb777/date/period"
)

func main() {
	g := &Generator{Created: map[string]bool{}, SymTypeMap: map[string]string{}}

	flag.StringVar(&g.BaseURL, "baseURL", "https://graph.microsoft.com/v1.0", "Base URL")
	flag.StringVar(&g.In, "in", "metadata/v1.0.xml", "Input file name")
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
