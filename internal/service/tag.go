package service

import (
	"github.com/itachilee/ginblog/internal/model"
	"github.com/itachilee/ginblog/pkg/app"
)

type CountTagRequest struct {
	Name  string `from:"name" binding:"max=100"`
	State uint   `from:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `from:"name" binding:"max=100"`
	State uint   `from:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name  string `from:"name" binding:"required,min=3,max=100"`
	State uint   `from:"state,default=1" binding:"oneof=0 1"`
}
type UpdateTagRequest struct {
	ID    uint   `from:"id" binding:"required,gte=1"`
	Name  string `from:"name" binding:"required,min=3,max=100"`
	State uint   `from:"state,default=1" binding:"oneof=0 1"`
}

type DeleteTagRequest struct {
	ID uint `from:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(params *CountTagRequest) (int64, error) {
	return svc.dao.CountTag(params.Name, params.State)
}

func (svc *Service) GetTagList(params *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(params.Name, params.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(params *CreateTagRequest) error {
	return svc.dao.CreateTag(params.Name, params.State)
}
func (svc *Service) UpdateTag(params *UpdateTagRequest) error {
	return svc.dao.UpdateTag(params.ID, params.Name, params.State)
}

func (svc *Service) DeleteTag(params *DeleteTagRequest) error {
	return svc.dao.DeleteTag(params.ID)
}
