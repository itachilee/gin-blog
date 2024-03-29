package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name  string `json:"name"`
	State uint   `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	err := db.Model(&t).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset > 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)

	}
	db = db.Where("state = ?", t.State)
	if err = db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(&Tag{}).Where("id = ?", t.ID).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ?", t.ID).Delete(&t).Error
}
