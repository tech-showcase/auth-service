package global

import (
	"github.com/tech-showcase/auth-service/helper"
	"github.com/tech-showcase/auth-service/model"
)

var UsersRepo model.UserRepo
var OAuth2Helper helper.OAuth2Helper

func init() {
	UsersRepo = model.NewUsersRepo()
	OAuth2Helper = helper.NewOAuth2Helper()
}
