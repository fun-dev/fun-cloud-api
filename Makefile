build-dev:
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
up-container:
	docker-compose -f deployments/docker-compose.dev.yml build container
	docker-compose -f deployments/docker-compose.dev.yml up container