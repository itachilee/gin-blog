package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/itachilee/ginblog/internal/service"
	"github.com/itachilee/ginblog/pkg/app"
	"github.com/itachilee/ginblog/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	params := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&params)

	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	var token string
	token, err = app.GenerateToken(params.Id, params.Username)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
