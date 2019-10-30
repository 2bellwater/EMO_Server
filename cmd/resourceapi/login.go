package resourceapi

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func SetLoginAPI(router *gin.Engine, db *gorm.DB) {

	//app_key := "a6387b7811a594082baa53208343832d"
	//redirect_uri := "http://15.165.60.87:8080/login/kakao"
//"https://kauth.kakao.com/oauth/authorize?client_id=547cc89df7b9c1ef14915757d9bc5948&redirect_uri=http://jsserver.iptime.org:8080/login/kakao&response_type=code",

	router.GET("/login/kakao", func(context *gin.Context) {
		exec.Command("xdg-open", "https://kauth.kakao.com/oauth/authorize?client_id=547cc89df7b9c1ef14915757d9bc5948&redirect_uri=http://jsserver.iptime.org:8080/login/kakao&response_type=code").Start()
		resp, err := http.Get("https://kauth.kakao.com/oauth/authorize?client_id=547cc89df7b9c1ef14915757d9bc5948&redirect_uri=http://jsserver.iptime.org:8080/login/kakao&response_type=code")
		if err != nil{
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(body)

		context.String(http.StatusOK, "logining...")

	})


	router.GET("/login/kakao_redirect", func(context *gin.Context) {
		context.String(http.StatusOK, "kako login")
	})


	router.GET("/login", func(context *gin.Context) {
		context.String(http.StatusOK, "kako login")
	})
}
