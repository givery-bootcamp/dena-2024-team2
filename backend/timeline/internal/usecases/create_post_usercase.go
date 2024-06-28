package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type CreatePostUsecase struct {
	createPostRepository interfaces.CreatePostRepository
	serverRepository     interfaces.ServerRepository
	channelRepository    interfaces.ChannelRepository
}

func NewCreatePostUsecase(cpr interfaces.CreatePostRepository, sr interfaces.ServerRepository, cr interfaces.ChannelRepository) *CreatePostUsecase {
	return &CreatePostUsecase{
		createPostRepository: cpr,
		serverRepository:     sr,
		channelRepository:    cr,
	}
}

func (u *CreatePostUsecase) Execute(userID int, serverID int, channelID int, content string) (*entities.Post, error) {
	// 存在しているサーバーか確認する
	if _, err := u.serverRepository.Get(serverID); err != nil {
		return nil, err
	}

	if _, err := u.channelRepository.Get(channelID); err != nil {
		return nil, err
	}

	return u.createPostRepository.Post(userID, channelID, content)
}
