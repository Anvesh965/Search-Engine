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
swag:

	@echo generating swagger docs.json

	swag init -d ./cmd/ -o ./docs --parseDependency

	@echo done with swagger docs
test:

	go test ./pkg/Controllers -v -cover

cover:

	go test -coverprofile coverage.out ./pkg/Controllers

	go tool cover -html coverage.out

