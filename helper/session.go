package helper

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

type (
	sessionStorage struct {
		engine *sessions.CookieStore
	}
	SessionStorage interface {
		Read(r *http.Request, name string) (result map[interface{}]interface{}, err error)
		Write(r *http.Request, w http.ResponseWriter, name string, value map[interface{}]interface{}) (err error)
	}

	sessionData struct {
		storage  SessionStorage
		name     string
		request  *http.Request
		response http.ResponseWriter
		data     map[interface{}]interface{}
	}
	SessionData interface {
		Get(key string) (result interface{}, ok bool)
		Set(key string, value interface{})
		Remove(key string)
		Save() (err error)
	}
)

var SessionStorageInstance SessionStorage

func NewSessionStorage() SessionStorage {
	var instance sessionStorage
	instance.engine = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	return &instance
}

func (instance *sessionStorage) Read(r *http.Request, name string) (result map[interface{}]interface{}, err error) {
	session, err := instance.engine.Get(r, name)
	if err != nil {
		return
	}

	result = session.Values
	return
}

func (instance *sessionStorage) Write(r *http.Request, w http.ResponseWriter, name string, value map[interface{}]interface{}) (err error) {
	session, err := instance.engine.Get(r, name)
	if err != nil {
		return
	}

	session.Values = value
	err = session.Save(r, w)
	if err != nil {
		return
	}

	return
}

func NewSessionData(storage SessionStorage, name string, r *http.Request, w http.ResponseWriter) (SessionData, error) {
	var instance sessionData
	var err error
	instance.storage = storage
	instance.name = name
	instance.request = r
	instance.response = w
	instance.data, err = storage.Read(r, name)

	return &instance, err
}

func (instance *sessionData) Get(key string) (result interface{}, ok bool) {
	result, ok = instance.data[key]
	return
}

func (instance *sessionData) Set(key string, value interface{}) {
	instance.data[key] = value
	return
}

func (instance *sessionData) Remove(key string) {
	delete(instance.data, key)
	return
}

func (instance *sessionData) Save() (err error) {
	err = instance.storage.Write(instance.request, instance.response, instance.name, instance.data)
	return
}
