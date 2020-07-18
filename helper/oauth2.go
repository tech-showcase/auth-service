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
	OAuth2Helper interface{
		HandleAuthorizeRequest(w http.ResponseWriter, r *http.Request) (err error)
		HandleTokenRequest(w http.ResponseWriter, r *http.Request) (err error)
	}
)

var OAuth2HelperInstance OAuth2Helper

func init() {
	OAuth2HelperInstance = NewOAuth2Helper()
}

func NewOAuth2Helper() OAuth2Helper {
	var instance oauth2Helper
	instance.server = initOAuth2Server()

	return &instance
}

func initOAuth2Server() *server.Server {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return srv
}

func (instance *oauth2Helper) HandleAuthorizeRequest(w http.ResponseWriter, r *http.Request) (err error) {
	err = instance.server.HandleAuthorizeRequest(w, r)
	return
}

func (instance *oauth2Helper) HandleTokenRequest(w http.ResponseWriter, r *http.Request) (err error) {
	err = instance.server.HandleTokenRequest(w, r)
	return
}
