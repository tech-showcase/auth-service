package cmd

import (
	"flag"
	"github.com/tech-showcase/auth-service/presenter"
)

func Parse() (args presenter.Args) {
	flag.IntVar(&args.Port, "port", 8080, "Port which service will listen to")
	flag.Parse()

	return
}
