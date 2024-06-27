package controllers

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetServers(ctx *gin.Context) {
	repository := repositories.NewServersRepository(DB(ctx))
	servers, err := repository.Get()
	if err != nil {
		handleError(ctx, 500, err)
	} else if servers != nil {
		ctx.JSON(200, serversConvertToJson(servers))
	} else {
		handleError(ctx, 404, errors.New("Not found"))
	}
}

func serversConvertToJson(servers []entities.Server) entities.Servers {
	jsonServers := make([]entities.Server, len(servers))
	for i, v := range servers {
		jsonServers[i] = entities.Server{
			Id:        v.Id,
			OwnerId:   v.OwnerId,
			Name:      v.Name,
			Icon:      v.Icon,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}
	return entities.Servers{Servers: jsonServers}
}
