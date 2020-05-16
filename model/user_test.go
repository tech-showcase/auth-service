package model

import (
	"reflect"
	"testing"
)

func TestUsers_AddOrUpdateUser(t *testing.T) {
	dummyInput := getDummyUser()
	expectedOutput := dummyInput

	mockUsers := NewUsersRepo()
	mockUsers.AddOrUpdateUser(dummyInput)

	userData := mockUsers.GetUserByPhone(expectedOutput.Phone)

	if userData.Phone == "" {
		t.Fatal("unexpected output")
	} else if !reflect.DeepEqual(userData, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func getDummyUser() User {
	dummyUser := User{
		UserData: UserData{
			Name:  "dummyName",
			Phone: "082222222222",
			Email: "dummy@email.com",
			Role:  "user",
		},
		UserCredential: UserCredential{
			Password: "1234",
		},
	}

	return dummyUser
}
