package global

import "github.com/tech-showcase/auth-service/model"

var UsersRepo model.UsersInterface

func init() {
	UsersRepo = model.NewUsersRepo()
}
