package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type GetChannelsRepository struct {
	Conn *gorm.DB
}

type Channels []*Channel
type Channel struct {
	Id        int
	ServerId  int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewChannelsRepository(conn *gorm.DB) *GetChannelsRepository {
	return &GetChannelsRepository{
		Conn: conn,
	}
}

func (r *GetChannelsRepository) Get() {
	obj :=[]Channel{}
	result := r.Conn.Find(&obj)
	fmt.Printf("result: %+v\n", obj)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		}
	}
}
