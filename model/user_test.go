package model

import (
	"reflect"
	"testing"
)

func TestUsers_AddUser(t *testing.T) {
	dummyInput := getDummyUser()
	expectedOutput := dummyInput

	mockUsers := NewUsersRepo()
	_, err := mockUsers.AddUser(dummyInput)

	userData := mockUsers.GetUserByPhone(expectedOutput.Phone)

	if err != nil {
		t.Fatal("an error occurred")
	} else if userData.Phone == "" {
		t.Fatal("unexpected output")
	} else if !reflect.DeepEqual(userData, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func TestUsers_UpdateUser(t *testing.T) {
	dummyInput := getDummyUser()
	expectedOutput := dummyInput

	mockUsers := NewUsersRepo()
	_, _ = mockUsers.AddUser(User{UserData: UserData{Phone: dummyInput.Phone}})
	err := mockUsers.UpdateUser(dummyInput)

	userData := mockUsers.GetUserByPhone(expectedOutput.Phone)

	if err != nil {
		t.Fatal("an error occurred")
	} else if userData.Phone == "" {
		t.Fatal("unexpected output")
	} else if !reflect.DeepEqual(userData, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func getDummyUser() User {
	dummyUser := User{
		UserData: UserData{
			Username: "dummyName",
			Phone:    "082222222222",
			Email:    "dummy@email.com",
		},
		UserCredential: UserCredential{
			Password: "1234",
		},
	}

	return dummyUser
}
