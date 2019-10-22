package resourceapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetHelloAPI(router *gin.Engine){

	router.GET("/hello/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello")
	})
}