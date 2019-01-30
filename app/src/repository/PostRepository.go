package repositories

import (
    "model"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostRepository struct {
    Db *gorm.DB
}

type Tag struct {
    Id int
}

func (pr *PostRepository) Create(title string, text string, category int, tagIds []int) {
    post := model.Post{
        Title:      title,
        Text:       text,
        CategoryID: category,
    }
    pr.Db.Create(&post).Scan(&post)

    for _, value := range tagIds {
        pr.Db.Exec("INSERT INTO post_tags (post_id, tag_id) VALUES (?,?)", post.ID, value)
    }
}

func (pr *PostRepository) GetAll() (db *gorm.DB) {
    var post model.Post
    return pr.Db.Preload("Category").Preload("Tags").Last(&post)
}