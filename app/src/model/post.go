package model

import (
    "github.com/jinzhu/gorm"
      _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Post struct {
    gorm.Model
    CategoryID int
    Category   Category
    Title      string `gorm:"size:255"`
    Text       string
    Tags       []Tag  `gorm:"many2many:post_tags;"`
}