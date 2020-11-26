package helper

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"log"
	"net/http"
)

type (
	oauth2Helper struct {
		server *server.Server
	}
	OAuth2Helper interface {
		HandleAuthorizeRequest(w http.ResponseWriter, r *http.Request) (err error)
		HandleTokenRequest(w http.ResponseWriter, r *http.Request) (err error)
	}
)

const (
	OAuth2SessionName = "oauth2-session-name"

	BaseOAuth2Url    = "/oauth2"
	LoginUrl         = BaseOAuth2Url + "/login"
	AuthorizationUrl = BaseOAuth2Url + "/authorization"

	BaseStaticFilepath  = "static"
	LoginStaticFilepath = BaseStaticFilepath + "/login.html"
	AuthStaticFilepath  = BaseStaticFilepath + "/authorization.html"
)

var OAuth2HelperInstance OAuth2Helper

func NewOAuth2Helper() OAuth2Helper {
	var instance oauth2Helper
	instance.server = initOAuth2Server()

	return &instance
}

func initOAuth2Server() *server.Server {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientStore.Set("082222333444", &models.Client{
		ID:     "082222333444",
		Secret: "082222333444",
		Domain: "http://localhost:8082/oauth2/token",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetUserAuthorizationHandler(AuthorizeUser)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return srv
}

func AuthorizeUser(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	sessionData, err := NewSessionData(SessionStorageInstance, OAuth2SessionName, r, w)
	if err != nil {
		return
	}

	userIDInf, _ := sessionData.Get("LoggedInUserID")
	userID = userIDInf.(string)
	sessionData.Set("AuthorizedUserID", userID)
	err = sessionData.Save()

	return
}

func (instance *oauth2Helper) HandleAuthorizeRequest(w http.ResponseWriter, r *http.Request) (err error) {
	err = instance.server.HandleAuthorizeRequest(w, r)
	return
}

func (instance *oauth2Helper) HandleTokenRequest(w http.ResponseWriter, r *http.Request) (err error) {
	err = instance.server.HandleTokenRequest(w, r)
	return
}
