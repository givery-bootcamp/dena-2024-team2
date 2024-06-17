package repositories

import (
	"errors"
	"fmt"
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type CreatePostRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type NewPost struct {
	Id        int
	ChannelId int
	UserId    int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewCreatePostRepository(conn *gorm.DB) *CreatePostRepository {
	return &CreatePostRepository{
		Conn: conn,
	}
}

func (r *CreatePostRepository) Post(post entities.Post) (*entities.Post, error) {
	result := r.Conn.Omit("DeletedAt").Create(&post)

	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", post)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return convertCreatePostRepositoryModelToEntity(&post), nil
}

func convertCreatePostRepositoryModelToEntity(v *entities.Post) *entities.Post {
	return &entities.Post{
		Id:        v.Id,
		ChannelId: v.ChannelId,
		UserId:    v.UserId,
		Content:   v.Content,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		DeletedAt: v.DeletedAt,
	}
}
