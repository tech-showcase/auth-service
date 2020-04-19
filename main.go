package main

import (
	"github.com/tech-showcase/auth-service/api"
	"github.com/tech-showcase/auth-service/cmd"
)

func main() {
	args := cmd.Parse()

	api.Activate(args.Port)
}
