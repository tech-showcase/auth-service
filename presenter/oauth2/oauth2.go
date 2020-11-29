package oauth2

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-showcase/auth-service/controller/oauth2"
	"github.com/tech-showcase/auth-service/helper"
	"net/http"
	"os"
)

func GetAuthorization(ctx *gin.Context) {
	sessionData, err := helper.NewSessionData(helper.SessionStorageInstance, helper.OAuth2SessionName, ctx.Request, ctx.Writer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	/*TODO:
	This function still doesn't work properly.
	After opening new browser's tab or another uri, even when user has already authorized, it still render authorization page.
	It will works properly for the second request and the next.
	*/
	if !oauth2.IsAuthorized(sessionData) {
		renderHTML(ctx.Writer, ctx.Request, helper.AuthStaticFilepath)
	} else {
		Authorize(ctx)
	}

	return
}

func Authorize(ctx *gin.Context) {
	sessionData, err := helper.NewSessionData(helper.SessionStorageInstance, helper.OAuth2SessionName, ctx.Request, ctx.Writer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	err = oauth2.Authorize(ctx.Writer, ctx.Request, sessionData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
}

func Token(ctx *gin.Context) {
	sessionData, err := helper.NewSessionData(helper.SessionStorageInstance, helper.OAuth2SessionName, ctx.Request, ctx.Writer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	sessionData.Get("AuthorizedUserID")

	err = oauth2.Token(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
}

func GetLogin(ctx *gin.Context) {
	sessionData, err := helper.NewSessionData(helper.SessionStorageInstance, helper.OAuth2SessionName, ctx.Request, ctx.Writer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	/*TODO:
	This function still doesn't work properly.
	After opening new browser's tab or another uri, even when user has already authorized, it still render login page.
	It will works properly for the second request and the next.
	*/
	if !oauth2.IsLoggedIn(sessionData) {
		renderHTML(ctx.Writer, ctx.Request, helper.LoginStaticFilepath)
	} else {
		ctx.Redirect(http.StatusFound, helper.AuthorizationUrl)
	}

	return
}

func PostLogin(ctx *gin.Context) {
	sessionData, err := helper.NewSessionData(helper.SessionStorageInstance, helper.OAuth2SessionName, ctx.Request, ctx.Writer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	if ctx.Request.Form == nil {
		if err = ctx.Request.ParseForm(); err != nil {
			return
		}
	}
	username := ctx.Request.Form.Get("username")

	err = oauth2.PostLogin(username, sessionData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	ctx.Redirect(http.StatusFound, helper.AuthorizationUrl)
}

func renderHTML(w http.ResponseWriter, r *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, r, file.Name(), fi.ModTime(), file)
}
