package resourceapi

import (
	"github.com/gin-gonic/gin"
	"github.com/waker_server/pkg/model"
	"net/http"
)


func SetHelloAPI(router *gin.Engine){

	router.GET("/hello/", func(context *gin.Context) {
		m := model.Hellomodel{
			Name: "LEEJONGSOO",
			Age:  29,
		}
		context.JSON(http.StatusOK,m)
	})

	router.POST("/hello/", func(context *gin.Context) {
		mhellomodel := &model.Hellomodel{}
		err := context.BindJSON(mhellomodel)
		if err != nil{
			context.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		context.JSON(http.StatusOK, mhellomodel)
	})
}