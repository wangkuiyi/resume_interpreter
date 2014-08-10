package main

import (
	"encoding/json"
	"flag"
	"fmt"
	peacock "github.com/wangkuiyi/peacock/inference_server/rpc"
	"github.com/wangkuiyi/resume_interpreter/rrlt"
	"io"
	"log"
	"net/rpc"
	"os"
)

func main() {
	input := flag.String("input", "./testdata/input", "The input JSON file")
	pserver := flag.String("peacock", "", "Address of Peacock server")
	flag.Parse()

	lsa, e := rpc.DialHTTP("tcp", *pserver)
	if e != nil {
		log.Fatal("Dialing Peacock server:", e)
	}

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

		for i, _ := range r.Data.Works {
			if len(r.Data.Works[i].JD) > 0 {
				var dist peacock.TopicDist
				if e := lsa.Call("Sampler.Interpret",
					&peacock.InferenceRequest{
						Text:      r.Data.Works[i].JD,
						NumReturn: 10},
					&dist); e != nil {
					log.Print("Sampler.Interpret error:", e)
				}

				descs := make(map[int32]string)
				if e := lsa.Call("Sampler.DescribeTopics", dist.Topics,
					&descs); e != nil {
					log.Print("Sampler.DescribeTopics error:", e)
				}

				fmt.Printf("%s\n", r.Data.Works[i].JD)
				for i, topic := range dist.Topics {
					fmt.Printf("  %f\t%s\n", dist.Weights[i], descs[topic])
				}
			}
		}
	}
}
