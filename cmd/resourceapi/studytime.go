package resourceapi

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/waker_server/pkg/database"
	"net/http"
)

func SetStudyTimeUpdate(router *gin.Engine, db *gorm.DB){

	router.POST("/timer/update", func(context *gin.Context){

		user := database.User{}

		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		userInfo := database.User{}

		db.First(&userInfo,"user_id = ?",user.UserID)
		db.Model(&userInfo).Update("total_study_time",userInfo.TotalStudyTime + user.TodayStudyTime)
		db.Model(&userInfo).Update("today_study_time",userInfo.TodayStudyTime + user.TodayStudyTime)

		context.JSON(http.StatusOK,gin.H{"status" : "success"})
	})

	//router.GET("timer/")
}
