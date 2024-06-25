package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type CreateChannelUsecase struct {
	serverRepository        interfaces.ServerRepository
	createChannelRepository interfaces.CreateChannelRepository
}

func NewCreateChannelUsecase(sr interfaces.ServerRepository, ccr interfaces.CreateChannelRepository) *CreateChannelUsecase {
	return &CreateChannelUsecase{
		serverRepository:        sr,
		createChannelRepository: ccr,
	}
}

func (uc *CreateChannelUsecase) Execute(serverId int, name string) (*entities.Channel, error) {
	// 存在しているサーバーか確認する
	if _, err := uc.serverRepository.Get(serverId); err != nil {
		return nil, err
	}

	result, err := uc.createChannelRepository.Post(serverId, name)
	return result, err
}
