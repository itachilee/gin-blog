package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/itachilee/ginblog/internal/model"
	"github.com/itachilee/ginblog/internal/service"
	"github.com/itachilee/ginblog/pkg/app"
	"github.com/itachilee/ginblog/pkg/convert"
	"github.com/itachilee/ginblog/pkg/errcode"
	"net/http"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// Get @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enum(0,1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Success 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) Get(c *gin.Context) {
	c.JSON(http.StatusOK, model.Tag{Name: c.Param("id")})
}

func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.NewPager(c)
	totalRows, err := svc.CountTag(&service.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCountTagFailed)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetTagListFailed)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&service.CreateTagRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateTagFailed)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&service.UpdateTagRequest{
		ID:    param.ID,
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateTagFailed)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt(),
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteTagFailed)
		return
	}
	response.ToResponse(gin.H{})
	return
}
