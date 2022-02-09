package models

import "log"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth Auth) TableName() string {
	return "blog_auth"
}

func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	// result := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if db == nil {
		log.Fatal("数据库异常")
	}
	result := db.Last(&auth)
	if result.Error != nil {
		return false, result.Error
	}

	if auth.ID > 0 {
		return true, nil
	}
	return false, nil
}
