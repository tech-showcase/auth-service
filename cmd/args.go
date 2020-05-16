package cmd

import (
	"flag"
)

type (
	Args struct {
		Port int `json:"port"`
	}
)

func Parse() (args Args) {
	flag.IntVar(&args.Port, "port", 8080, "Port which service will listen to")
	flag.Parse()

	return
}
