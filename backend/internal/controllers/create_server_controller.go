package controllers

import (
	"net/http"
	"time"

	"myapp/internal/entities"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

type createServerRequestParams struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func CreateServer(ctx *gin.Context) {
	var req createServerRequestParams
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	uid := getUserIdFromContext(ctx)
	uc := usecases.NewCreateServerUsecase()
	server, err := uc.Execute(ctx, req.Name, req.Icon, uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(200, convertToServerJson(server))
}

func convertToServerJson(server *entities.Server) ServerJson {
	return ServerJson{
		Id:        server.Id,
		OwnerId:   server.OwnerId,
		Name:      server.Name,
		Icon:      server.Icon,
		CreatedAt: server.CreatedAt,
		UpdatedAt: server.UpdatedAt,
	}
}

type ServerJson struct {
	Id        int       `json:"id"`
	OwnerId   int       `json:"owner_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
