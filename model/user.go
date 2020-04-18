package model

import "go-simple-app/presenter"

type (
	User struct {
		presenter.RegisterRequestStruct
		presenter.RegisterResponseStruct
	}
	Users map[string]User
)
