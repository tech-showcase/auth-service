package oauth2

import (
	"github.com/tech-showcase/auth-service/global"
	"net/http"
)

func Authorize(w http.ResponseWriter, r *http.Request) (err error) {
	err = global.OAuth2Helper.HandleAuthorizeRequest(w, r)
	return
}

func Token(w http.ResponseWriter, r *http.Request) (err error) {
	err = global.OAuth2Helper.HandleTokenRequest(w, r)
	return
}
