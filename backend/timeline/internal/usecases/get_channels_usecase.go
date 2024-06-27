package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type GetChannelsUsecase struct {
	serverRepository      interfaces.ServerRepository
	getChannelsRepository interfaces.GetChannelsRepository
}

func NewGetChannelsUsecase(sr interfaces.ServerRepository, gcr interfaces.GetChannelsRepository) *GetChannelsUsecase {
	return &GetChannelsUsecase{
		serverRepository:      sr,
		getChannelsRepository: gcr,
	}
}

func (uc *GetChannelsUsecase) Execute(serverId int) ([]entities.Channel, error) {
	// 存在しているサーバーか確認する
	if _, err := uc.serverRepository.Get(serverId); err != nil {
		return nil, err
	}

	result, err := uc.getChannelsRepository.Get(serverId)
	return result, err
}
