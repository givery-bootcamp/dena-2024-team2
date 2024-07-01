package repositories

import (
	"time"

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
		Table("posts").
		Where("id = ?", postId).
		Update("deleted_at", gorm.DeletedAt{Time: time.Now(), Valid: true}).
		Error; err != nil {
		return err
	}

	return nil
}
