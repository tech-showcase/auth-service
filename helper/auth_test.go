package helper

import (
	"github.com/tech-showcase/auth-service/model"
	"reflect"
	"testing"
)

func TestAuthBlueprint_GenerateToken(t *testing.T) {
	dummyPrivateClaims := getDummyClaims()
	dummyKey := getDummyKey()
	expectedToken := getDummyToken()

	authHelper := NewAuthHelper()
	token, err := authHelper.GenerateToken(dummyPrivateClaims, dummyKey)

	if err != nil {
		t.Fatal("error has occurred")
	} else if token != expectedToken {
		t.Fatal("unexpected output")
	}
}

func TestAuthBlueprint_ValidateToken(t *testing.T) {
	dummyKey := getDummyKey()
	dummyToken := getDummyToken()

	authHelper := NewAuthHelper()
	isValid, err := authHelper.ValidateToken(dummyToken, dummyKey)

	if err != nil {
		t.Fatal("error has occurred")
	} else if !isValid {
		t.Fatal("unexpected output")
	}
}

func TestAuthBlueprint_ParseToken(t *testing.T) {
	dummyToken := getDummyToken()
	expectedPrivateClaims := getDummyClaims()

	authHelper := NewAuthHelper()
	privateClaims, err := authHelper.ParseToken(dummyToken)

	if err != nil {
		t.Fatal("error has occurred")
	} else if !reflect.DeepEqual(privateClaims, expectedPrivateClaims) {
		t.Fatal("unexpected output")
	}
}

func getDummyToken() string {
	dummyToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImR1bW15TmFtZSIsInBob25lIjoiMDgyMTIzNDU2Nzg5IiwiZW1haWwiOiJkdW1teUBlbWFpbC5jb20iLCJpc19hY3RpdmUiOnRydWUsInRpbWVzdGFtcCI6MTI1Nzg5NDAwMH0.8FBX1VB7n_XZmw86fmbqfAvBKqVyz3Nx2rP0TcOcUlY"
	return dummyToken
}

func getDummyKey() string {
	dummyKey := "1234"
	return dummyKey
}

func getDummyClaims() PrivateClaims {
	userData := model.UserData{
		Username: "dummyName",
		Phone:    "082123456789",
		Email:    "dummy@email.com",
		IsActive: true,
	}
	dummyPrivateClaims := PrivateClaims{
		UserData:  userData,
		Timestamp: 1257894000,
	}
	return dummyPrivateClaims
}
