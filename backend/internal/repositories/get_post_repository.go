package repositories

import (
	"errors"
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type GetPostsRepository struct {
	Conn *gorm.DB
}

type Post struct {
	Id        int
	Title     string
	Body      string
	UserId    int
	UserName  string
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
	posts := []entities.Post{}
	result := r.Conn.Table("posts").Select("posts.id, posts.title, posts.body, posts.created_at, posts.updated_at,users.id, users.name").Joins("join users on posts.user_id = users.id").Scan(&posts)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return posts, nil
}
