package resourceapi

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/waker_server/pkg/database"
	"net/http"
)

func SetFaceBookLoginAPI(router *gin.Engine, db *gorm.DB) {

	var testmsg struct{
		UserID 		string 	`json:"userid"`
		UserToken	string	`json:"usertoken"`
	}

	router.GET("/login/facebook", func(context *gin.Context) {

		if err := context.ShouldBindJSON(&testmsg); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}

		user := database.User{}


		if err := db.Find(&user,"UserID=?",testmsg.UserID); err != nil{
			return
		}

		context.JSON(http.StatusOK, gin.H{"status":"success"})
	})


	router.POST("/login/facebook", func(context *gin.Context) {

		if err := context.ShouldBindJSON(&testmsg); err != nil{
			context.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
			return
		}

		user := database.User{
			UserID:testmsg.UserID,
			UserToken:testmsg.UserToken,
		}

		//curl -i -X GET "https://graph.facebook.com/20531316728?access_token={access-token}"
		//var nodeURL = "https://graph.facebook.com/2578019708953479?access_token="+user.UserToken;
		//resp, err := http.Get(nodeURL)
		//if err != nil {
		//	panic(err)
		//}
		//
		//respbody, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	panic(err)
		//}
		//
		//respbodyjson := json.NewDecoder(resp.Body).Decode()
		//fmt.Print(respbodyjson)
		//
		//defer resp.Body.Close()
		//
		//fmt.Println(respbody)



		if isNew := db.NewRecord(&user); isNew == true {
			//new account is registered
			db.Create(&user)
		}



		context.JSON(http.StatusOK,gin.H{"status" : "success"})
	})
}