package repositories

import (
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type GetPostsRepository struct {
	Conn *gorm.DB
}

type Post struct {
	Id        int       `gorm:"column:id"`
	ChannelId int       `gorm:"column:channel_id"`
	UserId    int       `gorm:"column:user_id"`
	UserName  string    `gorm:"column:name"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func NewGetPostRepository(conn *gorm.DB) *GetPostsRepository {
	return &GetPostsRepository{
		Conn: conn,
	}
}

func (r *GetPostsRepository) Get(channelId int) (entities.Posts, error) {
	posts := []Post{}
	if err := r.Conn.Table("posts").
		Select("posts.id, posts.channel_id, posts.user_id, users.name, posts.content, posts.created_at, posts.updated_at, posts.deleted_at").
		Joins("join users on posts.user_id = users.id").
		Where("channel_id = ?", channelId).
		Order("posts.created_at ASC").
		Scan(&posts).Error; err != nil {
		return nil, err
	}

	return convertToEntities(posts), nil
}

func convertToEntities(posts []Post) entities.Posts {
	entityPosts := make(entities.Posts, len(posts))
	for i, v := range posts {
		entityPosts[i] = entities.Post{
			Id:        v.Id,
			ChannelId: v.ChannelId,
			User:      entities.User{Id: v.UserId, Name: v.UserName},
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}
	return entityPosts
}
