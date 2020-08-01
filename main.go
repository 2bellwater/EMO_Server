package main

import (
	"github.com/EMO_Server/cmd/auth"
	"github.com/EMO_Server/cmd/configs"
	"github.com/EMO_Server/cmd/databases"
	"github.com/EMO_Server/cmd/webframework"
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

  //http router run
	router.Run()
}
