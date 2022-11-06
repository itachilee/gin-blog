package models // `User` 属于 `Company`，`CompanyID` 是外键
import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

//func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
//	m := tx.Statement.Dest.(map[string]interface{})
//	m["UpdatedAt"] = time.Now()
//	tx.Statement.Dest = m
//
//	return
//}

//func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
//	m := tx.Statement.Dest.(map[string]interface{})
//	m["DeletedAt"] = time.Now()
//	tx.Statement.Dest = m
//
//	return
//}
