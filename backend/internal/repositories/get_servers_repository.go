package repositories

import (
	"fmt"
	"log"
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type GetServersRepository struct {
	Conn *gorm.DB
}

type Servers []*Server
type Server struct {
	Id        int
	OwnerId   int
	Name      string
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewServersRepository(conn *gorm.DB) *GetServersRepository {
	return &GetServersRepository{
		Conn: conn,
	}
}

func (r *GetServersRepository) Get() ([]entities.Server, error) {
	obj := []Server{}
	if err := r.Conn.Find(&obj).Error; err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	fmt.Printf("result: %+v\n", obj)
	return convertServersRepositoryModelToEntity(obj), nil
}

func convertServersRepositoryModelToEntity(servers []Server) []entities.Server {
	entityServers := make([]entities.Server, len(servers))
	for i, v := range servers {
		entityServers[i] = entities.Server{
			Id:        v.Id,
			OwnerId:   v.OwnerId,
			Name:      v.Name,
			Icon:      v.Icon,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}
	}
	return entityServers
}
