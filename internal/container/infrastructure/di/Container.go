package di

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/adapters/controller"
	"github.com/fun-dev/fun-cloud-api/internal/container/adapters/gateway"
	"github.com/fun-dev/fun-cloud-api/internal/container/application/usecase"
	"github.com/fun-dev/fun-cloud-api/internal/container/infrastructure/server"
	"github.com/fun-dev/fun-cloud-api/pkg/kubernetes"
	"github.com/fun-dev/fun-cloud-api/pkg/redis"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()
	// --- Redis Driver --- //
	err := c.Provide(redis.NewRedisDriver)
	if err != nil {
		return nil, err
	}
	// --- Kubernetes Provider --- //
	err = c.Provide(kubernetes.NewKubernetesProvider)
	if err != nil {
		return nil, err
	}
	// --- Container Controller --- //
	err = c.Provide(controller.NewContainerController)
	if err != nil {
		return nil, err
	}
	// --- Container Delete Interactor --- //
	err = c.Provide(usecase.NewContainerDeleteInteractor)
	if err != nil {
		return nil, err
	}
	// --- Container Gateway --- //
	err = c.Provide(gateway.NewContainerGateway)
	if err != nil {
		return nil, err
	}
	// --- Auth Gateway --- //
	err = c.Provide(gateway.NewAuthGateway)
	if err != nil {
		return nil, err
	}
	// --- Server Driver --- //
	err = c.Provide(server.NewGinDriver)
	if err != nil {
		return nil, err
	}
	return c, nil
}
