package resourceapi

import (
	"github.com/EMO_Server/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func SetStudyTimeUpdate(router *gin.Engine, db *gorm.DB) {

	router.POST("/timer/update", func(context *gin.Context) {

		updateUserAccount := database.UserAccount{}

		if err := context.ShouldBindJSON(&updateUserAccount); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		currentUserAccount := database.UserAccount{}

		db.First(&currentUserAccount, "user_id = ?", updateUserAccount.UserID)
		db.Model(&currentUserAccount).Update("total_study_time", currentUserAccount.TotalStudyTime+updateUserAccount.TodayStudyTime)
		db.Model(&currentUserAccount).Update("today_study_time", currentUserAccount.TodayStudyTime+updateUserAccount.TodayStudyTime)

		context.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	//router.GET("timer/")
}
