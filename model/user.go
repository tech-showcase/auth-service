package model

import "github.com/tech-showcase/auth-service/presenter"

type (
	User struct {
		presenter.RegisterRequestStruct
		presenter.RegisterResponseStruct
	}
	Users map[string]User
)
