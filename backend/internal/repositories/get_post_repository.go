package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type GetPostsRepository struct {
	Conn *gorm.DB
}

type GetPosts struct {
	Posts []*Post
}

type Posts []*Post

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

func (r *GetPostsRepository) Get() {
	posts :=[]Post{}
	result := r.Conn.Find(&posts)
	fmt.Printf("result: %+v\n", result)
	if result.Error != nil {
		fmt.Printf("hoge")
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		}
	}
}
