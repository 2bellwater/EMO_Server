package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/waker_server/cmd/resourceapi"

	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
)

func SetRestAPI(router *gin.Engine, db *gorm.DB) {
	resourceapi.SetHelloAPI(router, db)
	resourceapi.SetLoginAPI(router, db)
}
