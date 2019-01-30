package model

import (
    "github.com/jinzhu/gorm"
      _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Tag struct {
    gorm.Model
    Name  string  `gorm:"size:255"`
    Posts []*Post `gorm:"many2many:post_tags;"`
}