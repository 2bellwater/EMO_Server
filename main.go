package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waker_server/cmd"

	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/redis"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// _, err := redisClient.Ping().Result()
	// if err != nil {
	// 	panic(err)
	// }

	db, err := sql.Open("mysql", "root:jspassword@(127.0.0.1:3306)/jsdatabase")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO jstable(userid, userpw) VALUES ('jongsooid3', 'jongsoopw3')")
	if err != nil {
		panic(err)
	}
	var id string
	err = db.QueryRow("SELECT userid FROM jstable").Scan(&id)
	if err != nil {
		panic(err)
	}
	println("id=", id)

	router := gin.Default()

	cmd.SetRestAPI(router, db)
	router.Run(":8080")

}
