package oauth2

import (
	"github.com/tech-showcase/auth-service/helper"
	"net/http"
)

func Authorize(w http.ResponseWriter, r *http.Request) (err error) {
	err = helper.OAuth2HelperInstance.HandleAuthorizeRequest(w, r)
	return
}

func Token(w http.ResponseWriter, r *http.Request) (err error) {
	err = helper.OAuth2HelperInstance.HandleTokenRequest(w, r)
	return
}
