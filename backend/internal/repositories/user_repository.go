package repositories

import (
	"errors"
	"log"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

type result struct {
	Id *uint
}

func (r *UserRepository) Authenticate(userName string, password string) (*uint, error) {
	result := &result{}
	if err := r.Conn.Select("user_id").Where("name = ? AND password=?", userName, password).Take(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entities.ErrNotFound
		}
		log.Printf("%v", err)
		return nil, err
	}
	return result.Id, nil
}
