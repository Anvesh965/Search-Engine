SHELL=cmd.exe

SEARCH_BINARY=Dockerfile


## up: starting app and mongo container with docker compose 

up:

	@echo Starting Docker Compose...

	docker-compose down 

	docker-compose up --build

	@echo Docker images started!

## down: stop docker compose

down:

	@echo Stopping docker compose...

	docker-compose down

	@echo Done!