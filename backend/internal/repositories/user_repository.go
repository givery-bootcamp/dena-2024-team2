package repositories

import (
	"errors"
	"log"
	"myapp/internal/entities"
	"myapp/internal/infrastructures"

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
	Id       uint
	Password string
}

func (r *UserRepository) Authenticate(userName string, password string) (uint, error) {
	result := &authResult{}
	if err := r.Conn.Table("users").Select("id, password").Where("name = ?", userName).Take(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, entities.ErrNotFound
		}
		log.Printf("%v", err)
		return 0, err
	}

	err := infrastructures.CompareHashAndPassword(result.Password, password)
	if err != nil {
		return 0, err
	}
	return result.Id, nil
}
