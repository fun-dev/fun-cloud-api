build:
	sudo docker-compose build
up:
	docker-compose up -d
logs:
	docker-compose logs -f
down:
	docker-compose down -v