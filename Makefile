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
	@CompileDaemon -build='make build' -graceful-kill -command='./out/api'

build:
	@CGO_ENABLED=0 go build -o out/api

deps:
	${call go, mod vendor}

test:
	${call go, test -v ./...}

#---- docker enviroment ----
ifdef DOCKER_COMPOSE_EXISTS
define go
	@docker-compose exec ${DOCKER_APP_CONTAINER} go ${1}
endef
endif