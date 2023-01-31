SHELL=cmd.exe

SEARCH_BINARY=Dockerfile

dirr := ./pkg/DatabaseConn
# mainDir := ./cmd/main.go

# coverage: $(dirr)
# 	go test $(dirr) -cover

# test: $(dirr)
# 	go test $(dirr)
# test-v: $(dirr)
# 	go test $(dirr) -v

# run: $(mainDir)
# 	go run $(mainDir)

# html:
# 	rm -f coverage.out
# 	go test -coverprofile=coverage.out $(dirr)
# 	go tool cover -html=coverage.out 
# 	rm -r coverage.out

test:

	go test ./... -v -cover

cover:

	go test -coverprofile coverage.out ./...

	go tool cover -html coverage.out

run:
	go run ./cmd/main.go


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



