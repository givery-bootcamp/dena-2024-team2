package repositories

import (
	"errors"
	"log"
	"myapp/internal/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerRepository struct {
	Conn *gorm.DB
}

func NewGetServerRepository(conn *gorm.DB) *ServerRepository {
	return &ServerRepository{
		Conn: conn,
	}
}

func (r *ServerRepository) Get(serverId int) (*entities.Server, error) {
	server := &entities.Server{}
	if err := r.Conn.First(&server, serverId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entities.ErrNotFound
		}
		log.Printf("%v", err)
		return nil, err
	}
	return server, nil
}

func (r *ServerRepository) Create(ctx *gin.Context, name string, icon string, uid int) (*entities.Server, error) {
	server := entities.Server{OwnerId: uid, Name: name, Icon: icon}
	result := r.Conn.Omit("DeletedAt").Create(&server)
	if result.Error != nil {
		return nil, result.Error
	}

	return &server, nil
}
