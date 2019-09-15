// +build ignore

package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	data := struct {
		BaseURL string
		Out     string
	}{}

	flag.StringVar(&data.BaseURL, "baseURL", "https://graph.microsoft.com/v1.0", "Base URL")
	flag.StringVar(&data.Out, "out", "metadata-v1.0.xml", "Output file name")
	flag.Parse()

	url := data.BaseURL + "/$metadata"
	log.Printf("Downloading %s to %s", url, data.Out)

	var err error
	var outFile io.WriteCloser
	if data.Out == "-" {
		outFile = os.Stdout
	} else {
		outFile, err = os.OpenFile(data.Out, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
		if err != nil {
			log.Printf("%s already exists, skipping", data.Out)
			return
		}
	}
	defer outFile.Close()

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	defer res.Body.Close()

	_, err = io.Copy(outFile, res.Body)
	if err != nil {
		log.Fatal(err)
	}
}
