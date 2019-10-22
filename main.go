package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waker_server/cmd"
)

func main() {

	router := gin.Default()

	cmd.SetRestAPI(router)
	router.Run(":8080")

}