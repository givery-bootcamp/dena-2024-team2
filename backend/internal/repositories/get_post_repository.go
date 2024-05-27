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
	posts := []Post{}
	result := r.Conn.Table("posts").Select("posts.id, posts.title, posts.body, posts.user_id, posts.created_at, posts.updated_at, users.name").Joins("join users on posts.user_id = users.id").Scan(&posts)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return convertGetPostsRepositoryModelToEntity(posts), nil
}

func convertGetPostsRepositoryModelToEntity(posts []Post) entities.Posts {
	entityPosts := make([]entities.Post, len(posts))
	for i, v := range posts {
		entityPosts[i] = entities.Post{
			Id:        v.Id,
			Title:     v.Title,
			Body:      v.Body,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}
	}
	return entityPosts
}
