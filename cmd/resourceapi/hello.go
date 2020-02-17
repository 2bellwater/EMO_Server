package resourceapi

import (
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waker_server/pkg/model"
)

func SetHelloAPI(router *gin.Engine, db *gorm.DB) {

	router.GET("/hello/", func(context *gin.Context) {
		m := model.Hellomodel{
			Name:     "Ahyun",
			Favorite: "Potato",
		}
		context.JSON(http.StatusOK, m)
	})

	router.POST("/hello/", func(context *gin.Context) {
		mhellomodel := &model.Hellomodel{}
		err := context.BindJSON(mhellomodel)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, mhellomodel)
	})

}
