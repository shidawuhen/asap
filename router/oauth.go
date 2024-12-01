package router

import (
	"asap/controller/oauth"
	"github.com/gin-gonic/gin"
)

func oauthFunc(router *gin.Engine) {
	router.GET("/oauth/jwt", oauth.JWT)
}
