package main

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/adapters/controller"
	"github.com/fun-dev/fun-cloud-api/internal/container/adapters/gateway"
	"github.com/fun-dev/fun-cloud-api/internal/container/application/usecase"
	"github.com/fun-dev/fun-cloud-api/internal/container/infrastructure/server"
	"github.com/fun-dev/fun-cloud-api/pkg/kubernetes"
	"github.com/fun-dev/fun-cloud-api/pkg/mongo"
	"log"
)

var (
	_k8sProvider   = kubernetes.NewKubernetesProvider()
	_mongoDriver   = mongo.NewMongoDriver()
	_containerRepo = gateway.NewContainerGateway(_k8sProvider, _mongoDriver)
	_authRepo      = gateway.NewAuthGateway()
	_read          = usecase.NewContainerReadInteractor(_containerRepo, _authRepo)
	_create        = usecase.NewContainerCreateInteractor(_containerRepo, _authRepo)
	_delete        = usecase.NewContainerDeleteInteractor(_containerRepo, _authRepo)
	_ctrl          = controller.NewContainerController(_read, _create, _delete)
	serverDriver   = server.NewGinDriver(_ctrl)
)

func init() {
	if err := _k8sProvider.InitKubectl("", ""); err != nil {
		log.Fatal(err)
	}
	if err := _k8sProvider.InitK8SClient(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// --- initialize program --- //
	log.Printf("info: listing ...")
	if err := serverDriver.Run(); err != nil {
		log.Fatal(err)
	}
}
