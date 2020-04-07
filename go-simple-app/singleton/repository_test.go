package singleton

import (
	"go-simple-app/model"
	"go-simple-app/presenter"
	"reflect"
	"testing"
)

func TestUsersRepoBlueprint_AddOrUpdateUser(t *testing.T) {
	dummyInput := model.User{
		RegisterRequestStruct: presenter.RegisterRequestStruct{
			Name:  "dummyName",
			Phone: "082222222222",
			Role:  "user",
		},
		RegisterResponseStruct: presenter.RegisterResponseStruct{
			Password: "1234",
		},
	}
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
