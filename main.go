package main

import (
	"fmt"
	"github.com/tech-showcase/auth-service/api"
	"github.com/tech-showcase/auth-service/cmd"
)

func main() {
	fmt.Println("Hi, I am Auth Service!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
