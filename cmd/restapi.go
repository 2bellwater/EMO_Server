package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/waker_server/resourceapi"
)

func SetRestAPI(router *gin.Engine) {
	resourceapi.SetHelloAPI(router)

}
