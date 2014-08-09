package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/wangkuiyi/resume_indexer/rrlt"
	"io"
	"log"
	"os"
)

func main() {
	input := flag.String("input", "./testdata/input", "The input JSON file")
	flag.Parse()

	i, e := os.Open(*input)
	if e != nil {
		log.Fatal("Cannot open input file:", *input)
	}

	d := json.NewDecoder(i)
	for {
		var r rrlt.Record
		if e := d.Decode(&r); e == io.EOF {
			break
		} else if e != nil {
			log.Fatal("Failed parsing resume:", e)
		}

		fmt.Printf("%+v", r)
	}
}
