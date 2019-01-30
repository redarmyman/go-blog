package main

import (
    "os"
    "fmt"
    "model"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
    db, err := gorm.Open("postgres", "host=db user=postgres dbname=blog password=example sslmode=disable")
    if err != nil {
        fmt.Printf("%v\n",err)
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&model.Category{}, &model.Post{}, &model.Tag{})

    os.Exit(0)
}