package model

type (
	Users map[string]User
	User  struct {
		UserData
		UserCredential
	}
	UserData struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}
	UserCredential struct {
		Password string `json:"password"`
	}
)
