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
		GetUserByPhone(phone string) model.User
	}
)

func NewUsersRepo() UsersRepoInterface {
	var instance UsersRepoBlueprint

	return &instance
}

func (instance *UsersRepoBlueprint) AddOrUpdateUser(user model.User) {
	instance.Data[user.Phone] = user
}

func (instance *UsersRepoBlueprint) GetUserByPhone(phone string) model.User {
	if user, ok := instance.Data[phone]; ok {
		return user
	}

	return model.User{}
}
