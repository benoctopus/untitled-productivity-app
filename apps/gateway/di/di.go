package di

import (
	"fmt"
	"gateway/config"
	"proto"

	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Tasks      proto.TaskServiceClient
	Tags       proto.TagServiceClient
	BaseClient *grpc.ClientConn
}

var c *dig.Container

func Init() (func(), error) {
	c = dig.New()

	// TODO: secure connection

	err := c.Provide(func() *config.Config {
		return config.New()
	})

	if err != nil {
		return nil, err
	}
	err = c.Provide(func(conf *config.Config) (*ServiceClient, error) {
		gc, err := grpc.Dial(conf.TaskServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}

		return &ServiceClient{
			Tasks:      proto.NewTaskServiceClient(gc),
			Tags:       proto.NewTagServiceClient(gc),
			BaseClient: gc,
		}, nil
	})

	if err != nil {
		return nil, err
	}

	return func() {
		c.Invoke(func(sc *ServiceClient) {
			sc.BaseClient.Close()
		})
	}, nil

}

func Invoke[t *ServiceClient | *config.Config | any](fn func(t) error) error {
	if c == nil {
		panic(fmt.Errorf("DI container is not initialized"))
	}

	return c.Invoke(fn)
}
