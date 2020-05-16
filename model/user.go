package model

type (
	User struct {
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

	users          map[string]User
	UsersInterface interface {
		AddOrUpdateUser(user User)
		GetUserByPhone(phone string) User
	}
)

func NewUsersRepo() UsersInterface {
	instance := make(users)

	return &instance
}

func (instance users) AddOrUpdateUser(user User) {
	instance[user.Phone] = user
}

func (instance users) GetUserByPhone(phone string) User {
	if user, ok := instance[phone]; ok {
		return user
	}

	return User{}
}
