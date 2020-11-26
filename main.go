package main

import (
	"fmt"
	"github.com/tech-showcase/auth-service/api"
	"github.com/tech-showcase/auth-service/cmd"
	"github.com/tech-showcase/auth-service/config"
	"github.com/tech-showcase/auth-service/helper"
)

func init() {
	config.Instance = config.Read()

	helper.OAuth2HelperInstance = helper.NewOAuth2Helper()
}

func main() {
	fmt.Println("Hi, I am Auth Service!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
