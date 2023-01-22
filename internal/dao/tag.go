package dao

import (
	"github.com/itachilee/ginblog/internal/model"
	"github.com/itachilee/ginblog/pkg/app"
)

func (d *Dao) CountTag(name string, state uint) (int64, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint) error {
	tag := model.Tag{Name: name, State: state}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint, name string, state uint) error {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	tag.ID = id
	values := map[string]interface{}{
		"state": state,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}
func (d *Dao) DeleteTag(id uint) error {
	tag := model.Tag{}
	tag.ID = id
	return tag.Delete(d.engine)
}
