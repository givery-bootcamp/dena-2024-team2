package usecases

import "myapp/internal/interfaces"

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

func (u *GetChannelsUsecase) Execute() {
}
