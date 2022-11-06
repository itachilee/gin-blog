package api

import (
	"fmt"
	"github.com/itachilee/ginblog/models"
	"github.com/itachilee/ginblog/pkg/e"
	"github.com/itachilee/ginblog/pkg/logging"
	"github.com/itachilee/ginblog/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `vaild:"Required;MaxSize(50)"`
	Password string `vaild:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Printf("sadad %s%s", username, password)
	vaild := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := vaild.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist, id, err := models.CheckAuth(username, password)
		if err != nil {
			fmt.Printf("%v", err)
		}
		if isExist {
			token, err := util.GenerateToken(id, username)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}

	} else {
		for _, err := range vaild.Errors {
			logging.Info(err.Key, err.Message)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
