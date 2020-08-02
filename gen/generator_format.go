package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"runtime"
	"sync"

	"golang.org/x/tools/imports"
)

func (g *Generator) Format() error {
	var wgErr error
	wg := &sync.WaitGroup{}
	q := make(chan struct{}, runtime.NumCPU())
	for path := range g.Created {
		if wgErr != nil {
			break
		}
		q <- struct{}{}
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			log.Printf("Formatting %s", path)
			err := func() error {
				inB, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				outB, err := imports.Process(path, inB, nil)
				if err != nil {
					return err
				}
				if bytes.Equal(inB, outB) {
					return nil
				}
				return ioutil.WriteFile(path, outB, 0644)
			}()
			if err != nil && wgErr == nil {
				wgErr = err
			}
			<-q
		}(path)
	}
	wg.Wait()
	return wgErr
}
