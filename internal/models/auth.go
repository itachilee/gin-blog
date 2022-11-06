package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// func (auth Auth) TableName() string {
// 	return "blog_auth"
// }

func CheckAuth(username, password string) (bool, int, error) {
	var auth Auth
	result := DB.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if result.Error != nil {
		return false, 0, result.Error
	}

	if auth.ID > 0 {
		return true, auth.ID, nil
	}
	return false, 0, nil
}
