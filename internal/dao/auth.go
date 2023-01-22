package dao

import "github.com/itachilee/ginblog/internal/model"

func (d *Dao) GetAuth(id int, username string) (model.Auth, error) {
	auth := model.Auth{Username: username}
	auth.ID = uint(id)
	return auth.Get(d.engine)
}
