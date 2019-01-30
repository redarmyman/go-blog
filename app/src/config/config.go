package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "repository"
)

func Config() (*repositories.CategoryRepository, *repositories.TagRepository, *repositories.PostRepository) {
    return &repositories.CategoryRepository{Db: db()}, &repositories.TagRepository{Db: db()}, &repositories.PostRepository{Db: db()}
}

func db() (*gorm.DB) {
    db, err := gorm.Open("postgres", "host=db user=postgres dbname=blog password=example sslmode=disable")
    if err != nil {
        panic("cannot connect with database")
    }

    return db
}