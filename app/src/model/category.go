package model

import (
    "github.com/jinzhu/gorm"
      _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Category struct {
    gorm.Model
    Name  string  `gorm:"type:varchar(100);unique;not null"`
    Posts []Post  `gorm:"foreignkey:CategoryID"`
}