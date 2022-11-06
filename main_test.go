package main

import (
	"fmt"
	models2 "github.com/itachilee/ginblog/internal/models"
	"github.com/itachilee/ginblog/internal/setting"
	"testing"
)

func init() {

	setting.Setup()
	models2.InitDb()
}

func Test_belongs_to(t *testing.T) {

	user := models2.User{
		Name:      "test1",
		CompanyID: 1,
		Company:   models2.Company{ID: 1, Name: "test"},
	}
	models2.DB.Create(&user)
}

func Test_query_belongs_to(t *testing.T) {

	user := models2.User{}
	models2.DB.First(&user)
	fmt.Println(user)
	fmt.Println(user.Company)

}

func Test_preload_belongs_to(t *testing.T) {

	user := models2.User{}
	models2.DB.Preload("Company").First(&user)
	fmt.Println(user)
	fmt.Println(user.Company)

}

func TestUpdateHook(t *testing.T) {

	user := models2.User{}
	models2.DB.First(&user)

	fmt.Println(user.ID)
	models2.DB.Model(&user).Updates(map[string]interface{}{"name": "afterUpdated  hhh "})

}
