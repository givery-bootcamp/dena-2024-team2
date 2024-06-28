package repositories

import (
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type DeletePostsRepository struct {
	Conn *gorm.DB
}

func NewDeletePostRepository(conn *gorm.DB) *DeletePostsRepository {
	return &DeletePostsRepository{
		Conn: conn,
	}
}

func (r *DeletePostsRepository) Delete(postId int) error {
	if err := r.Conn.
		Delete(&entities.Post{}, postId).
		Error; err != nil {
		return err
	}

	return nil
}
