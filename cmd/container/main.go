package container

import (
	"github.com/fun-dev/fun-cloud-api/internal/container/infrastructure/di"
	"github.com/fun-dev/fun-cloud-api/internal/container/infrastructure/server"
	"log"
)

func main() {
	// --- initialize program --- //
	c, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	err = c.Provide(func(server *server.GinDriver) {
		err := server.Run()
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
