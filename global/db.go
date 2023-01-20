package global

import (
	"github.com/itachilee/ginblog/internal/model"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
)

func SetupDBEngine() error {

	var err error
	DBEngine, err = model.NewDBEngine()
	if err != nil {
		return err
	}
	return nil
}
