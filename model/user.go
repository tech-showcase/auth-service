package model

import (
	"errors"
	"strconv"
)

type (
	User struct {
		UserData
		UserCredential
	}
	UserData struct {
		Username string `json:"username"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		IsActive bool   `json:"is_active"`
	}
	UserCredential struct {
		Password string `json:"password"`
	}

	userRepo struct {
		data   map[string]User
		lastID int
	}
	UserRepo interface {
		AddUser(user User) (id string, err error)
		UpdateUser(user User) (err error)
		GetUserByPhone(phone string) (user User)
	}
)

var UserRepoInstance UserRepo

func NewUsersRepo() UserRepo {
	instance := userRepo{}
	instance.data = make(map[string]User)

	return &instance
}

func (instance userRepo) AddUser(user User) (id string, err error) {
	if _, ok := instance.data[user.Phone]; !ok {
		instance.data[user.Phone] = user
		instance.lastID++
		return strconv.Itoa(instance.lastID), nil
	}

	return "", errors.New("user already exists")
}

func (instance userRepo) UpdateUser(user User) (err error) {
	if _, ok := instance.data[user.Phone]; ok {
		instance.data[user.Phone] = user
		return nil
	}

	return errors.New("user is not found")
}

func (instance userRepo) GetUserByPhone(phone string) User {
	if user, ok := instance.data[phone]; ok {
		return user
	}

	return User{}
}
