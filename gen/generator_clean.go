package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
