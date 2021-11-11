up:
	@docker-compose up

build:
	@docker-compose build

down:
	@docker-compose down
	@docker volume prune -f
	@docker network prune -f

.PHONY: up build down
