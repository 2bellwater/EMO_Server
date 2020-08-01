package cmd

import (
	"github.com/EMO_Server/cmd/resourceapi"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
)

func SetRestAPI(router *gin.Engine, db *gorm.DB) {

	// auth rest api
	resourceapi.SetHelloAPI(router, db)

	// db rest api
	resourceapi.SetStudyTimeUpdate(router, db)
}
