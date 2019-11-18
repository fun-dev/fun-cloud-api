package di

import (
	"github.com/fun-dev/ccms-poc/adapters/gateway"
	"github.com/fun-dev/ccms-poc/adapters/gateway/repository"
	"github.com/fun-dev/ccms-poc/application/usecase"
	"go.uber.org/dig"
)

func New() (*dig.Container, error) {
	c := dig.New()
	// --- Container Delete Interactor --- //
	err := c.Provide(usecase.NewContainerDeleteInteractor)
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
	return c, nil
}
