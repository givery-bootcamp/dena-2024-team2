package repositories

import (
	"myapp/internal/entities"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerRepository struct {
	Conn *gorm.DB
}

func NewServerReopsitory(conn *gorm.DB) *ServerRepository {
	return &ServerRepository{
		Conn: conn,
	}
}

type server struct {
	Id        int
	OwnerId   int
	Name      string
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (r *ServerRepository) Create(ctx *gin.Context, name string, icon string, uid int) (*entities.Server, error) {
	server := server{OwnerId: uid, Name: name, Icon: icon}
	result := r.Conn.Create(&server)
	if result.Error != nil {
		return nil, result.Error
	}

	return convertServerRepositoryModelToEntity(&server), nil
}

func convertServerRepositoryModelToEntity(s *server) *entities.Server {
	return &entities.Server{
		Id:        s.Id,
		OwnerId:   s.OwnerId,
		Name:      s.Name,
		Icon:      s.Icon,
		CreatedAt: s.CreatedAt,
	}
}
