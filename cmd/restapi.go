package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/waker_server/cmd/resourceapi"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
)

func SetRestAPI(router *gin.Engine, db *gorm.DB) {

	// auth rest api
	resourceapi.SetHelloAPI(router, db)
	resourceapi.SetKakaoLoginAPI(router, db)
	resourceapi.SetFaceBookLoginAPI(router,db)

	// db rest api
	resourceapi.SetStudyTimeUpdate(router, db)
}
