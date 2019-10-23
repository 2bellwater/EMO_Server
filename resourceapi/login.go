package resourceapi

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetLoginAPI(router *gin.Engine, db *sql.DB) {

	router.GET("/login/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello")
	})
}
