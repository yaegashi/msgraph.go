package main

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const Tab = 2

const (
	OpenTag = iota
	CloseTag
	Content
)

// XML dumb indentation routine
func indent(b []byte) []byte {
	o := &bytes.Buffer{}
	t := OpenTag
	i := 0
	p := 0
	for p < len(b) {
		if rune(b[p]) == '<' {
			if p+1 < len(b) && rune(b[p+1]) == '/' {
				t = CloseTag
				i -= Tab
			} else {
				t = OpenTag
			}
		} else {
			t = Content
		}
		for x := 0; x < i; x++ {
			o.WriteRune(' ')
		}
		if t == Content {
			for p < len(b) {
				if rune(b[p]) == '<' {
					break
				}
				o.WriteByte(b[p])
				p++
			}
		} else {
			for c := true; c && p < len(b); p++ {
				if rune(b[p]) == '>' {
					c = false
					if t == OpenTag && (rune(b[p-1]) != '/' && rune(b[p-1]) != '?') {
						i += Tab
					}
				}
				o.WriteByte(b[p])
			}
		}
		o.WriteRune('\n')
	}
	return o.Bytes()
}

func main() {
	data := struct {
		Pretty  bool
		BaseURL string
		Out     string
	}{}

	flag.BoolVar(&data.Pretty, "pretty", false, "Pretty output")
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
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if data.Pretty {
		b = indent(b)
	}

	_, err = outFile.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
