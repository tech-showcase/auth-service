package main

import (
	"flag"
	"go-simple-app/api"
)

func main() {
	portPtr := flag.Int("port", 8080, "Port which apps will listen to")
	flag.Parse()

	api.Activate(*portPtr)
}
