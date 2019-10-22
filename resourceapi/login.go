package resourceapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetLoginAPI(router *gin.Engine){

	router.GET("/hello/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello")
	})
}