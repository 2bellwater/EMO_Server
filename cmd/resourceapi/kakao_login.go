package resourceapi

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/waker_server/pkg/database"
	"io/ioutil"
	"net/http"
)

type t struct {
	NickName				string	`json:"nickName, string"`
	ProfileImageURL			string	`json:"profileImageURL, string"`
	ThumbnailURL			string	`json:"thumbnailURL, string"`
	CountryISO				string	`json:"countryISO, string"`
}

func SetKakaoLoginAPI(router *gin.Engine, db *gorm.DB) {

	router.POST("/login/kakao", func(context *gin.Context) {

		user := database.User{}

		if err := context.ShouldBindJSON(&user); err != nil{
			context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		profileURL := "https://kapi.kakao.com/v1/api/talk/profile?access_token="+user.UserToken
		resp, err := http.Get(profileURL)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Get kakao profile error : " + err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Read kakao profile error : " + err.Error())
		}

		err = json.Unmarshal(body, user.Kakaoprofile)
		if err != nil {
			fmt.Println("Unmarshal kakao profile error : " + err.Error())
		}

		if isNew := db.NewRecord(&user); isNew == true {
			fmt.Println("New kakao user created!")
			db.Create(&user)
		}

		context.JSON(http.StatusOK,gin.H{"status" : "success"})
	})

}
