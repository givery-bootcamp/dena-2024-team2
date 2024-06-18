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

type authResult struct {
	Id uint
}

func (r *UserRepository) Authenticate(userName string, password string) (uint, error) {
	result := &authResult{}
	if err := r.Conn.Table("users").Select("id").Where("name = ? AND password=?", userName, password).Take(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, entities.ErrNotFound
		}
		log.Printf("%v", err)
		return 0, err
	}
	return result.Id, nil
}
