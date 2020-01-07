package databases

import (
	"github.com/jinzhu/gorm"
	"github.com/waker_server/pkg/database"
)


func SetDBAutomigration(db *gorm.DB){

	db.AutoMigrate(&database.UserAccount{})
}