package database

type UserAccount struct {
	UserID			string			`json:"userid" gorm:"unique"`
	UserEmail 		string 			`json:"useremail"`
	UserToken		string 			`json:"usertoken"`
	TotalStudyTime	int64			`json:"totalstudytime"`
	TodayStudyTime	int64			`json:"todaystudytime"`
	//Profile			KakaoProfile
	KakaoProfile
}

type KakaoProfile struct {
	NickName				string	`json:"nickName"`
	ProfileImageURL			string	`json:"profileImageURL"`
	ThumbnailURL			string	`json:"thumbnailURL"`
	CountryISO				string	`json:"countryISO"`
}