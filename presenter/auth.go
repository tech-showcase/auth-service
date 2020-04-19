package presenter

import "github.com/tech-showcase/auth-service/model"

type (
	PrivateClaims struct {
		model.UserData
		Timestamp int64 `json:"timestamp"`
	}

	RegisterRequestStruct struct {
		model.UserData
	}

	RegisterResponseStruct struct {
		model.UserCredential
	}

	LoginRequestStruct struct {
		Phone string `json:"phone"`
		model.UserCredential
	}

	LoginResponseStruct struct {
		Token string `json:"token"`
	}

	AuthHeaderStruct struct {
		Authorization string `header:"Authorization"`
	}
)
