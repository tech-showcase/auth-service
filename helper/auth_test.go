package helper

import (
	"go-simple-app/presenter"
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
	claims, err := authObject.ParseToken(dummyToken, dummyKey)

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
	claims, err := authObject.ParseTokenWithoutKey(dummyToken)

	if err != nil {
		t.Fatal("error has occurred")
	} else if !reflect.DeepEqual(claims, expectedPrivateClaims) {
		t.Fatal("unexpected output")
	}
}

func getDummyToken() string {
	dummyToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZHVtbXlOYW1lIiwicGhvbmUiOiIwODIxMjM0NTY3ODkiLCJyb2xlIjoiZHVtbXlSb2xlIiwidGltZXN0YW1wIjoiMTI1Nzg5NDAwMCJ9.Govi5TIIF7oE1W8yKlV_t81B0-a-gJJ0T2GMO-8tB68"
	return dummyToken
}

func getDummyKey() string {
	dummyKey := "1234"
	return dummyKey
}

func getDummyClaims() presenter.PrivateClaims {
	dummyPrivateClaims := presenter.PrivateClaims{
		Name:      "dummyName",
		Phone:     "082123456789",
		Role:      "dummyRole",
		Timestamp: "1257894000",
	}
	return dummyPrivateClaims
}
