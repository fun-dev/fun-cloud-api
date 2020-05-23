package di

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/adapters/gateway"
	"github.com/fun-dev/fun-cloud-api/internal/container/adapters/service"
	"github.com/fun-dev/fun-cloud-api/internal/container/application/usecase"
	"github.com/fun-dev/fun-cloud-api/internal/container/infrastructure/server"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudk8s"
	"github.com/fun-dev/fun-cloud-api/pkg/cloudstore"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()
	// --- Redis Driver --- //
	err := c.Provide(cloudstore.NewRedisDriver)
	if err != nil {
		return nil, err
	}
	// --- Kubernetes Provider --- //
	err = c.Provide(cloudk8s.NewKubernetesProvider)
	if err != nil {
		return nil, err
	}
	// --- Container Controller --- //
	err = c.Provide(service.NewContainerController)
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
