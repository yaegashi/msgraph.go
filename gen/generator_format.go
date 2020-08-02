package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"golang.org/x/tools/imports"
)

func (g *Generator) Format() error {
	for path := range g.Created {
		log.Printf("Formatting %s", path)
		inB, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		outB, err := imports.Process(path, inB, nil)
		if err != nil {
			return err
		}
		if !bytes.Equal(inB, outB) {
			err = ioutil.WriteFile(path, outB, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
