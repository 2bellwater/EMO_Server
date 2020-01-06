package database

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID			string			`json:"userid" gorm:"type:varchar(100);unique;not null"`
	UserEmail 		string 			`json:"useremail" gorm:"unique;not null;type:varchar(100);"`
	UserToken		string 			`json:"usertoken" gorm:"type:varchar(500);"`
	TotalStudyTime	int64			`json:"totalstudytime" gorm:"type:int64;"`
	TodayStudyTime	int64			`json:"todaystudytime" gorm:"type:int64;"`
	Kakaoprofile 	KakaoProfile 	`json:"kakaoprofile"`
}

type KakaoProfile struct {
	NickName				string	`json:"nickName, string"`
	ProfileImageURL			string	`json:"profileImageURL, string"`
	ThumbnailURL			string	`json:"thumbnailURL, string"`
	CountryISO				string	`json:"countryISO, string"`
}