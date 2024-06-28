package usecase_test

import (
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/interfaces"
	"myapp/internal/usecases"
	"reflect"
	"testing"
)

type serverRepositoryMock struct {
	interfaces.ServerRepository
	FakeGet func(serverId int) (*entities.Server, error)
}

type getChannelsRepositoryMock struct {
	interfaces.GetChannelsRepository
	FakeGet func(serverId int) ([]entities.Channel, error)
}

func (m *serverRepositoryMock) Get(serverId int) (*entities.Server, error) {
	return m.FakeGet(serverId)
}

func (m *getChannelsRepositoryMock) Get(serverId int) ([]entities.Channel, error) {
	return m.FakeGet(serverId)
}
func Test_GetChannelUsecase_1(t *testing.T) {
	type serverFakes struct {
		Get func(serverId int) (*entities.Server, error)
	}
	type getChannelsFakes struct {
		Get func(serverId int) ([]entities.Channel, error)
	}
	type args struct {
		serverId int
	}

	tests := []struct {
		name             string
		serverFakes      serverFakes
		getChannelsFakes getChannelsFakes
		args             args
		want             []entities.Channel
		wantErr          bool
	}{
		{"get successfully",
			serverFakes{
				func(serverId int) (*entities.Server, error) {
					get := &entities.Server{Id: serverId, Name: "server1", OwnerId: 1, Icon: "icon"}
					return get, nil
				}},
			getChannelsFakes{
				func(serverId int) ([]entities.Channel, error) {
					get := []entities.Channel{
						{Id: 1, ServerId: serverId, Name: "channel1"},
					}
					return get, nil
				},
			},
			args{1},
			[]entities.Channel{
				{Id: 1, ServerId: 1, Name: "channel1"},
			},
			false,
		},
		{"server is not found",
			serverFakes{
				func(serverId int) (*entities.Server, error) {
					return nil, fmt.Errorf("err")
				}},
			getChannelsFakes{
				func(serverId int) ([]entities.Channel, error) {
					get := []entities.Channel{
						{Id: 1, ServerId: serverId, Name: "channel1"},
					}
					return get, nil
				},
			},
			args{1},
			[]entities.Channel{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sr := &serverRepositoryMock{}
			sr.FakeGet = tt.serverFakes.Get
			gcr := &getChannelsRepositoryMock{}
			gcr.FakeGet = tt.getChannelsFakes.Get
			uc := usecases.NewGetChannelsUsecase(sr, gcr)
			got, err := uc.Execute(tt.args.serverId)
			if (err != nil) && tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChannelsUsecase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
