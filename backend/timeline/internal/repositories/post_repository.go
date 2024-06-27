package repositories

import (
	"errors"
	"log"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

func NewPostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{
		Conn: conn,
	}
}

func (r *PostRepository) Get(postId int) (*entities.Post, error) {
	post := &entities.Post{}
	if err := r.Conn.First(&post, postId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entities.ErrNotFound
		}
		log.Printf("%v", err)
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) Update(postId int, content string) (*entities.Post, error) {
	var post post
	if err := r.Conn.Model(&post).Update("content", content).Error; err != nil {
		return nil, err
	}
	return convertCreatePostRepositoryModelToEntity(&post), nil
}
