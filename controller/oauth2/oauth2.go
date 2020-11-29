package oauth2

import (
	"github.com/tech-showcase/auth-service/helper"
	"net/http"
	"net/url"
)

func Authorize(w http.ResponseWriter, r *http.Request, sessionData helper.SessionData) (err error) {
	var form url.Values
	if v, ok := sessionData.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	r.Form = form

	err = helper.OAuth2HelperInstance.HandleAuthorizeRequest(w, r)
	return
}

func SaveAuthorizationData(r *http.Request, sessionData helper.SessionData) {
	_, ok := sessionData.Get("ReturnUri")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}
		sessionData.Set("ReturnUri", r.Form)
		sessionData.Save()
	}
}

func IsLoggedIn(sessionData helper.SessionData) (isLoggedIn bool) {
	_, isLoggedIn = sessionData.Get("LoggedInUserID")
	return
}

func IsAuthorized(sessionData helper.SessionData) (isAuthorized bool) {
	_, isAuthorized = sessionData.Get("AuthorizedUserID")
	return
}

func Token(w http.ResponseWriter, r *http.Request) (err error) {
	err = helper.OAuth2HelperInstance.HandleTokenRequest(w, r)
	return
}

func PostLogin(username string, sessionData helper.SessionData) (err error) {
	sessionData.Set("LoggedInUserID", username)

	err = sessionData.Save()
	if err != nil {
		return
	}

	return
}
