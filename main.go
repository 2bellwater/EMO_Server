package main

import (
	"github.com/waker_server/cmd/auth"
	"github.com/waker_server/cmd/configs"
	"github.com/waker_server/cmd/databases"
	"github.com/waker_server/cmd/webframework"
)

func main() {

	SGConfig := configs.GetSroundConfig()

	//init database
	var gormDB databases.GormDB
	var err error

	if SGConfig.DatabaseType == configs.SG_DB_MYSQL {
		gormDB, err = databases.NewGormDB(SGConfig.MySQLConfig)
		if err != nil {
			//TODO
		}
	}

	//http web framework
	var router *webframework.GinRouter
	router = webframework.NewGinWebFrameWork()

	//auth
	var authAPIs []auth.Auth
	for _, a := range SGConfig.LoginType {
		if a == configs.SG_AUTH_KAKAO {
			kakaoLoginApi := auth.NewKakaoLoginAPI(router.Engine, gormDB.DB)
			kakaoLoginApi.OnLoginAPI()
			authAPIs = append(authAPIs, kakaoLoginApi)
		}
		if a == configs.SG_AUTH_FACEBOOK {
			facebookLoginApi := auth.NewFaceBookLoginAPI(router.Engine, gormDB.DB)
			facebookLoginApi.OnLoginAPI()
			authAPIs = append(authAPIs, facebookLoginApi)
		}
	}

	router.Run()
}
