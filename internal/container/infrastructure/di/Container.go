package di

import (
	"github.com/fun-dev/ccms/internal/container/adapters/controller"
	"github.com/fun-dev/ccms/internal/container/adapters/gateway"
	"github.com/fun-dev/ccms/internal/container/application/usecase"
	provider2 "github.com/fun-dev/ccms/internal/container/infrastructure/provider"
	driver2 "github.com/fun-dev/ccms/internal/container/infrastructure/server"
	"github.com/fun-dev/ccms/pkg/redis"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()
	// --- Kubectl Driver --- //
	err := c.Provide(driver2.NewKubectlDriver)
	if err != nil {
		return nil, err
	}
	// --- Redis Driver --- //
	err = c.Provide(redis.NewRedisDriver)
	if err != nil {
		return nil, err
	}
	// --- Kubernetes Provider --- //
	err = c.Provide(provider2.NewKubernetesProvider)
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
	err = c.Provide(driver2.NewGinDriver)
	if err != nil {
		return nil, err
	}
	return c, nil
}
