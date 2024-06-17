package repositories

import (
	"gorm.io/gorm"
)

type SignUpRepository struct {
	Conn *gorm.DB
}

func NewSignUpRepository(conn *gorm.DB) *SignUpRepository {
	return &SignUpRepository{
		Conn: conn,
	}
}

func (r *SignUpRepository) SignUp() {

}
