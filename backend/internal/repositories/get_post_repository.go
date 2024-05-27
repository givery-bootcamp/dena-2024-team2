package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"myapp/internal/entities"
)

type GetPostsRepository struct {
	Conn *gorm.DB
}

type GetPosts struct {
	Posts []*Post
}

type Post struct {
	Id int
	Title string
	Body string
	UserId int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}



func NewGetPostRepository(conn *gorm.DB) *GetPostsRepository {
	return &GetPostsRepository{
		Conn: conn,
	}
}

func (r *GetPostsRepository) Get() (entities.Posts, error) {
	posts :=[]Post{}
	result := r.Conn.Find(&posts)
	if result.Error != nil {
		fmt.Printf("hoge")
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		}
	}
	return convertGetPostsRepositoryModelToEntity(posts), nil
}

func convertGetPostsRepositoryModelToEntity(posts []Post) entities.Posts {
  entityPosts:=make([]entities.Post, len(posts))
  for i, v := range posts {
	entityPosts[i] = entities.Post{
		Id: v.Id,
		Title: v.Title,
		Body: v.Body,
		UserId: v.UserId,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		DeletedAt: v.DeletedAt,
	}
  }
  return entityPosts
}
