package main

import (
	"fmt"
	"github.com/itachilee/ginblog/models"
	"github.com/itachilee/ginblog/pkg/setting"
	"testing"
)

func init() {

	setting.Setup()
	models.InitDb()
}

func Test_belongs_to(t *testing.T) {

	user := models.User{
		Name:      "test1",
		CompanyID: 1,
		Company:   models.Company{ID: 1, Name: "test"},
	}
	models.DB.Create(&user)
}

func Test_query_belongs_to(t *testing.T) {

	user := models.User{}
	models.DB.First(&user)
	fmt.Println(user)
	fmt.Println(user.Company)

}

func Test_preload_belongs_to(t *testing.T) {

	user := models.User{}
	models.DB.Preload("Company").First(&user)
	fmt.Println(user)
	fmt.Println(user.Company)

}

func TestUpdateHook(t *testing.T) {

	user := models.User{}
	models.DB.First(&user)

	fmt.Println(user.ID)
	models.DB.Model(&user).Updates(map[string]interface{}{"name": "afterUpdated  hhh "})

}
