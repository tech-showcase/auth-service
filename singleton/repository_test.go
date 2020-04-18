package singleton

import (
	"go-simple-app/model"
	"go-simple-app/presenter"
	"reflect"
	"testing"
)

func TestUsersRepoBlueprint_AddOrUpdateUser(t *testing.T) {
	dummyInput := getDummyUser()
	expectedOutput := dummyInput

	mockUsersRepo := UsersRepoBlueprint{
		Data: make(model.Users),
	}
	mockUsersRepo.AddOrUpdateUser(dummyInput)

	if userData, ok := mockUsersRepo.Data[expectedOutput.Phone]; !ok {
		t.Fatal("unexpected output")
	} else if !reflect.DeepEqual(userData, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func TestUsersRepoBlueprint_GetUserByPhone(t *testing.T) {
	dummyInput := getDummyUsers()
	expectedOutput := getOneDummyUserFromUsers(dummyInput)

	mockUsersRepo := UsersRepoBlueprint{
		Data: dummyInput,
	}
	user := mockUsersRepo.GetUserByPhone(expectedOutput.Phone)

	if user.Phone == "" {
		t.Fatal("unexpected output")
	} else if !reflect.DeepEqual(user, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func getDummyUser() model.User {
	dummyUser := model.User{
		RegisterRequestStruct: presenter.RegisterRequestStruct{
			Name:  "dummyName",
			Phone: "082222222222",
			Role:  "user",
		},
		RegisterResponseStruct: presenter.RegisterResponseStruct{
			Password: "1234",
		},
	}

	return dummyUser
}

func getDummyUsers() model.Users {
	dummyUser := getDummyUser()

	dummyUsers := make(model.Users)
	dummyUsers[dummyUser.Phone] = dummyUser

	return dummyUsers
}

func getOneDummyUserFromUsers(users model.Users) model.User {
	for _, user := range users {
		return user
	}

	return model.User{}
}