include .env
export

NAME := notion_service
NETWORK := notion_service
DOCKER := docker
DC := docker-compose

# REGISTRY=gitlab-registry.42paris.fr

# PROD MODE COMMANDS
all: up logs
reload: down all
re: clean all
up:
	$(DC) up --detach
logs:
	$(DC) logs --follow --tail 1000
build:
	$(DC) build
stop:
	$(DC) stop
down:
	$(DC) down
ps:
	$(DC) ps
clean:
	$(DC) down --rmi all --volumes
	rm -fr ./tmp
	rm -fr ./postgres-data
	rm -fr ./redisdata

# CHECK FORMAT BEFORE GENERATING API
validate:
	$(DOCKER) run --rm \
		--volume $(shell pwd):/go/src/gitlab.42paris.fr/notion_service \
		--workdir /go/src/gitlab.42paris.fr/notion_service \
		quay.io/goswagger/swagger validate ./gen/swagger.yml

# RE GENERATE API
gen:
	rm -frv ./pkg/models ./pkg/api/operations
	$(DOCKER) run --rm \
		--volume $(shell pwd):/go/src/gitlab.42paris.fr/notion_service \
		--workdir /go/src/gitlab.42paris.fr/notion_service \
		quay.io/goswagger/swagger generate server --name $(NAME) --spec ./gen/swagger.yml --target ./pkg --main-package=../../cmd/notion-server/

# DEV MODE COMMANDS
dev: 
	$(DC) -f ./docker-compose.dev.yml up --detach
devall: dev devlogs
devclean: 
	$(DC) -f ./docker-compose.dev.yml down --rmi all --volumes
	rm -fr ./tmp
	rm -fr ./postgres-data
	rm -fr ./redisdata
devlogs:
	$(DC) -f ./docker-compose.dev.yml logs --follow --tail 1000
devre: devclean devall
devstop: 
	$(DC) -f ./docker-compose.dev.yml stop
devdown: 
	$(DC) -f ./docker-compose.dev.yml down
devps: 
	$(DC) -f ./docker-compose.dev.yml ps

air.init:
	$(DOCKER) run --rm \
		--volume $(shell pwd):/go/src/gitlab.42paris.fr/notion_service \
		--workdir /go/src/gitlab.42paris.fr/notion_service \
		cosmtrek/air init


# NETWORK COMMANDS
network.create:
	$(DOCKER) network create $(NETWORK)
network.remove:
	$(DOCKER) network remove $(NETWORK)
