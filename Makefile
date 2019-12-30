build:
	docker-compose -f deployments/docker-compose.dev.yml build container
up:
	docker-compose up -d
logs:
	docker-compose logs -f
down:
	docker-compose -f deployments/docker-compose.dev.yml down -v
up-db:
	docker-compose -f deployments/docker-compose.dev.yml up db
up-auth:
	docker-compose -f deployments/docker-compose.dev.yml build auth
	docker-compose -f deployments/docker-compose.dev.yml up auth
clean:
	docker-compose -f deployments/docker-compose.dev.yml down -v
up-container:
	docker-compose -f deployments/docker-compose.dev.yml build container
	docker-compose -f deployments/docker-compose.dev.yml up -d container_mongo container_mongo_express container
	docker-compose -f deployments/docker-compose.dev.yml exec container sh ./build/container/dev_init.sh
	#docker-compose -f deployments/docker-compose.dev.yml logs -f container
develop-container:
	docker-compose -f deployments/docker-compose.dev.yml build container
	docker-compose -f deployments/docker-compose.dev.yml up -d container_mongo container_mongo_express container
	docker-compose -f deployments/docker-compose.dev.yml exec container ash
gen-mock:
	mockgen -source ./internal/container/domain/container/ContainerRepository.go -destination ./internal/container/application/usecase/repository_mock/ContainerRepository.go