DOCKER_APP_CONTAINER = app
DOCKER_COMPOSE_EXISTS := $(shell command -v docker-compose 2> /dev/null)

.PHONY: docker-up docker-stop docker-ssh migrate-up migrate-down run build deps test build-api-docs

docker-up:
	@docker-compose up -d

docker-stop:
	@docker-compose stop

docker-ssh:
	@docker-compose exec $(DOCKER_APP_CONTAINER) bash

run:
	@CompileDaemon -build='make build-local' -graceful-kill -command='./build/local'

run-prod:
	@CompileDaemon -build='make build' -graceful-kill -command='./build/prod'

build:
	env GOOS=linux GOARCH=amd64 go build -o build/prod -v cmd/lambda/main.go

build-local:
	env GOOS=linux GOARCH=amd64 go build -o build/local -v cmd/local/main.go


deps:
	${call go, mod vendor}

test:
	${call con, ./test lambda }

#---- docker enviroment ----
ifdef DOCKER_COMPOSE_EXISTS
define go
	@docker-compose exec ${DOCKER_APP_CONTAINER} go ${1}
endef
endif
ifdef DOCKER_COMPOSE_EXISTS
define con
	@docker-compose exec ${DOCKER_APP_CONTAINER} ${1}
endef
endif