package databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/waker_server/cmd/configs"
	"github.com/waker_server/pkg/database"
	"log"
)

type GormDB struct {
	DB *gorm.DB
}

func NewGormDB(config configs.MySQLConfig) (GormDB, error) {
	gormDB := GormDB{}
	var err error

	gormDB.DB, err = gorm.Open("mysql", config.User+":"+config.Password+"@("+config.Endpoint+")/"+config.DBname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("DB connection ERROR")
		log.Println(err.Error())
		return gormDB, err
	}

	SetDBAutomigration(gormDB.DB)
	return gormDB, nil
}

func SetDBAutomigration(db *gorm.DB) {
	db.AutoMigrate(&database.KakaoProfile{})
	db.AutoMigrate(&database.UserAccount{})
	db.AutoMigrate(&database.FriendsMatchingArray{})
}

func (g *GormDB) insert() {

}
func (g *GormDB) get() {

}
func (g *GormDB) delete() {

}
