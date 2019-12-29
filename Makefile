build:
	sudo docker-compose build
up:
	docker-compose up -d
logs:
	docker-compose logs -f
down:
	docker-compose -f deployments/docker-compose.dev.yml down -v
up-db:
	docker-compose -f deployments/docker-compose.dev.yml up db