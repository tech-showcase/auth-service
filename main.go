package main

import (
	"encoding/gob"
	"fmt"
	"github.com/tech-showcase/auth-service/api"
	"github.com/tech-showcase/auth-service/cmd"
	"github.com/tech-showcase/auth-service/config"
	"github.com/tech-showcase/auth-service/helper"
	"github.com/tech-showcase/auth-service/model"
	"net/url"
)

func init() {
	config.Instance = config.Read()

	model.UserRepoInstance = model.NewUsersRepo()

	helper.OAuth2HelperInstance = helper.NewOAuth2Helper()

	gob.Register(url.Values{})
	helper.SessionStorageInstance = helper.NewSessionStorage()
}

func main() {
	fmt.Println("Hi, I am Auth Service!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
