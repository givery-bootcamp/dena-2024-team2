package repositories

import (
	"myapp/internal/entities"
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

type reqPost struct {
	Id        int           `json:"id"`
	ChannelId int           `json:"channel_id"`
	UserID    int           `json:"user_id"`           // 外部キーを示すフィールドを追加
	User      entities.User `gorm:"foreignKey:UserID"` // 外部キーの関係を定義
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt *time.Time    `json:"deleted_at"`
}

func (reqPost) TableName() string {
	return "posts"
}

func (r *DeletePostsRepository) Delete(postId int) error {
	if err := r.Conn.
		Where("id = ?", postId).
		Delete(&reqPost{}).
		Error; err != nil {
		return err
	}

	return nil
}
