package usecase_test

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
	"myapp/internal/usecases"
	"testing"
)

type serverRepositoryMock struct {
	interfaces.ServerRepository
	FakeGet func(serverId int) (*entities.Server, error)
}

type getChannelsRepository struct {
	interfaces.GetChannelsRepository
	FakeGet func(serverId int) ([]entities.Channel, error)
}

func (m *serverRepositoryMock) Get(serverId int) (*entities.Server, error) {
	return m.FakeGet(serverId)
}

func (m *getChannelsRepository) Get(serverId int) ([]entities.Channel, error) {
	return m.FakeGet(serverId)
}

func Test_GetChannelUsecase_1(t *testing.T) {
	serverRepo := &serverRepositoryMock{
		FakeGet: func(serverId int) (*entities.Server, error) {
			get := &entities.Server{Id: serverId, Name: "server1", OwnerId: 1, Icon: "icon"}
			return get, nil
		},
	}

	getChannelsRepo := &getChannelsRepository{
		FakeGet: func(serverId int) ([]entities.Channel, error) {
			get := []entities.Channel{
				{Id: 1, ServerId: serverId, Name: "channel1"},
			}
			return get, nil
		},
	}

	uc := usecases.NewGetChannelsUsecase(serverRepo, getChannelsRepo)
	got, err := uc.Execute(1)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 1 {
		t.Errorf("usecase.Execute() should return length 1 array but got = %d length array", len(got))
	}
}
