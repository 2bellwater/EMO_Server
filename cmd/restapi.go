package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/waker_server/resourceapi"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func SetRestAPI(router *gin.Engine, db *sql.DB) {
	resourceapi.SetHelloAPI(router, db)

}
