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

	UsersInterface interface {
		AddOrUpdateUser(user User)
		GetUserByPhone(phone string) User
	}
)

func NewUsersRepo() UsersInterface {
	instance := make(Users)

	return &instance
}

func (instance Users) AddOrUpdateUser(user User) {
	instance[user.Phone] = user
}

func (instance Users) GetUserByPhone(phone string) User {
	if user, ok := instance[phone]; ok {
		return user
	}

	return User{}
}
