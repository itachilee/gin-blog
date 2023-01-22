package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model

	// public keys
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("id = ? AND username =?", a.ID, a.Username)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}
