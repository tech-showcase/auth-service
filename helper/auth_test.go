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

	authObject := NewAuthBlueprint()
	token, err := authObject.GenerateToken(dummyPrivateClaims, dummyKey)

	if err != nil {
		t.Fatal("error has occurred")
	} else if token != expectedToken {
		t.Fatal("unexpected output")
	}
}

func TestAuthBlueprint_ParseToken(t *testing.T) {
	dummyKey := getDummyKey()
	dummyToken := getDummyToken()
	expectedPrivateClaims := getDummyClaims()

	authObject := NewAuthBlueprint()
	claims, err := authObject.ParseAndValidateToken(dummyToken, dummyKey)

	if err != nil {
		t.Fatal("error has occurred")
	} else if !reflect.DeepEqual(claims, expectedPrivateClaims) {
		t.Fatal("unexpected output")
	}
}

func TestAuthBlueprint_ParseTokenWithoutKey(t *testing.T) {
	dummyToken := getDummyToken()
	expectedPrivateClaims := getDummyClaims()

	authObject := NewAuthBlueprint()
	claims, err := authObject.ParseToken(dummyToken)

	if err != nil {
		t.Fatal("error has occurred")
	} else if !reflect.DeepEqual(claims, expectedPrivateClaims) {
		t.Fatal("unexpected output")
	}
}

func getDummyToken() string {
	dummyToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZHVtbXlOYW1lIiwicGhvbmUiOiIwODIxMjM0NTY3ODkiLCJlbWFpbCI6ImR1bW15QGVtYWlsLmNvbSIsInJvbGUiOiJkdW1teVJvbGUiLCJ0aW1lc3RhbXAiOjEyNTc4OTQwMDB9.LnlGvHq5_T7QSBHY2F0lpkA6OM9YokK0afEk1-5tZmc"
	return dummyToken
}

func getDummyKey() string {
	dummyKey := "1234"
	return dummyKey
}

func getDummyClaims() PrivateClaims {
	userData := model.UserData{
		Username:  "dummyName",
		Phone: "082123456789",
		Email: "dummy@email.com",
	}
	dummyPrivateClaims := PrivateClaims{
		UserData:  userData,
		Timestamp: 1257894000,
	}
	return dummyPrivateClaims
}
