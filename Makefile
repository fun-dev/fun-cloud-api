up-store:
	docker-compose -f ./deployments/auth/docker-compose.yaml up -d fcp_auth_store
down:
	docker-compose -f ./deployments/auth/docker-compose.yaml down -v
gen-mock:
	mockgen -source ./internal/container/domain/container/ContainerRepository.go -destination ./internal/container/application/usecase/repository_mock/ContainerRepository.go