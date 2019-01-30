package repositories

import (
    "model"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type TagRepository struct {
    Db *gorm.DB
}

func (tr *TagRepository) Create(name string) {
    tag := model.Tag{Name: name}
    tr.Db.Create(&tag)
}

func (tr *TagRepository) GetAll() (db *gorm.DB) {
    var tags []model.Tag
    return tr.Db.Find(&tags)
}