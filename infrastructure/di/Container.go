package di

import (
	"github.com/fun-dev/ccms-poc/adapters/controller"
	"github.com/fun-dev/ccms-poc/adapters/gateway"
	"github.com/fun-dev/ccms-poc/application/usecase"
	"github.com/fun-dev/ccms-poc/infrastructure/driver"
	"github.com/fun-dev/ccms-poc/infrastructure/provider"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()
	// --- Kubectl Driver --- //
	err := c.Provide(driver.NewKubectlDriver)
	if err != nil {
		return nil, err
	}
	// --- Redis Driver --- //
	err = c.Provide(driver.NewRedisDriver)
	if err != nil {
		return nil, err
	}
	// --- Kubernetes Provider --- //
	err = c.Provide(provider.NewKubernetesProvider)
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
	err = c.Provide(driver.NewGinDriver)
	if err != nil {
		return nil, err
	}
	return c, nil
}
