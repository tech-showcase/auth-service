package singleton

import "go-simple-app/model"

var UsersRepo UsersRepoInterface

func init() {
	UsersRepo = NewUsersRepo()
}

type (
	UsersRepoBlueprint struct {
		Data model.Users
	}
	UsersRepoInterface interface {
		AddOrUpdateUser(user model.User)
	}
)

func NewUsersRepo() UsersRepoInterface {
	var instance UsersRepoBlueprint

	return &instance
}

func (instance *UsersRepoBlueprint) AddOrUpdateUser(user model.User) {
	instance.Data[user.Phone] = user
}
