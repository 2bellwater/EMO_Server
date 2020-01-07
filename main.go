package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/waker_server/cmd"
	"github.com/waker_server/cmd/databases"
)

func main() {

	db, err := gorm.Open("mysql", "root:jspassword@(127.0.0.1:3306)/jsdatabase?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println("DB connection ERROR")
		panic(err)
	}
	defer db.Close()

	router := gin.Default()

	databases.SetDBAutomigration(db)
	cmd.SetRestAPI(router, db)

	router.Run(":8080")

}