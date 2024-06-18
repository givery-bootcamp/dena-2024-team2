package repositories

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SignUpRepository struct {
	Conn *gorm.DB
}

type Account struct {
	Id        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewSignUpRepository(conn *gorm.DB) *SignUpRepository {
	return &SignUpRepository{
		Conn: conn,
	}
}

func (r *SignUpRepository) SignUp(name string, password string) error {

	account := Account{}
	r.Conn.Table("users").Select("name").Where("name = ?", name).Find(&account)
	if account.Name != "" {
		// name がすでに使われているときエラーにする
		return fmt.Errorf("%v", "その名前は他のアカウントで使用されています。")
	}

	// 新しいアカウントのとき保存する
	newAccount := Account{Name: name, Password: password}
	r.Conn.Table("users").Omit("DeletedAt").Create(&newAccount)

	return nil
}
