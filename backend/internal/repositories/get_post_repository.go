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
	Id        int
	ChannelId int
	UserId    int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewGetPostRepository(conn *gorm.DB) *GetPostsRepository {
	return &GetPostsRepository{
		Conn: conn,
	}
}

func (r *GetPostsRepository) Get(channelId int) (entities.Posts, error) {
	posts := []Post{}
	if err := r.Conn.Table("posts").
		Select("posts.id, posts.channel_id, posts.user_id, posts.content, posts.created_at, posts.updated_at").
		Joins("join users on posts.user_id = users.id").
		Where("channel_id = ?", channelId).Scan(&posts).Error; err != nil {
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
			UserId:    v.UserId,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}
	}
	return entityPosts
}
