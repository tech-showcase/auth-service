package global

import "github.com/tech-showcase/auth-service/model"

var UsersRepo model.UserRepo

func init() {
	UsersRepo = model.NewUsersRepo()
}
