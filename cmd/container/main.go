package container

import (
	"github.com/fun-dev/ccms/internal/container/infrastructure/di"
	"github.com/fun-dev/ccms/internal/container/infrastructure/driver"
	"log"
)

func main() {
	// --- initialize program --- //
	c, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	err = c.Provide(func(server *driver.GinDriver) {
		err := server.Run()
		if err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
