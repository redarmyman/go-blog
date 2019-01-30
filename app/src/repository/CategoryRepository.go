package repositories

import (
    "model"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type CategoryRepository struct {
   Db *gorm.DB
}

func (cr *CategoryRepository) Create(name string) {
   category := model.Category{Name: name}
   cr.Db.Create(&category)
}

func (cr *CategoryRepository) GetAll() (db *gorm.DB) {
   var categories []model.Category
   return cr.Db.Find(&categories)
}

