include .env
export

NAME := notion_service
NETWORK := notion_service
DOCKER := docker
DC := docker-compose

PROD_USER := zouz
PROD_HOST := notion-api.int.42paris.fr
PROD_PORT := 22
PROD_DC := dc-prod.yml
PROD_SSH := ssh://$(PROD_USER)@$(PROD_HOST):$(PROD_PORT)

DOCKER_PROD := $(DOCKER) --host $(PROD_SSH)
DC_PROD := $(DC) --host $(PROD_SSH) --file $(PROD_DC)

REGISTRY=gitlab-registry.42paris.fr
PROD_API_IMAGE := $(REGISTRY)/$(NAME):prod

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
network.create:
	$(DOCKER) network create $(NETWORK)
network.remove:
	$(DOCKER) network remove $(NETWORK)

###
### DEV 
###

# DEV COMMANDS
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
		--volume $(shell pwd):/go/src/gitlab.42paris.fr/utilities/notion_service \
		--workdir /go/src/gitlab.42paris.fr/utilities/notion_service \
		cosmtrek/air init

# CHECK FORMAT BEFORE GENERATING API
validate:
	$(DOCKER) run --rm \
		--volume $(shell pwd):/go/src/gitlab.42paris.fr/utilities/notion_service \
		--workdir /go/src/gitlab.42paris.fr/utilities/notion_service \
		quay.io/goswagger/swagger validate ./gen/swagger.yml

# RE GENERATE API
gen:
	rm -frv ./pkg/models ./pkg/api/operations
	$(DOCKER) run --rm \
		--volume $(shell pwd):/go/src/gitlab.42paris.fr/utilities/notion_service \
		--workdir /go/src/gitlab.42paris.fr/utilities/notion_service \
		quay.io/goswagger/swagger generate server --name $(NAME) --spec ./gen/swagger.yml --target ./pkg --main-package=../../cmd/notion-server/

###
### PRODUCTION 
###

prod: prod.up prod.logs
prod.build:
	$(DOCKER) login --password $(GITLAB_ACCESS_TOKEN) --username $(GITLAB_USERNAME) $(REGISTRY)
	$(DOCKER) build --build-arg GITLAB_ACCESS_TOKEN=$(GITLAB_ACCESS_TOKEN) --file ./deploy/Dockerfile.api-prod.build --tag $(PROD_API_IMAGE) .
	$(DOCKER) push $(PROD_API_IMAGE)
prod.up:
	$(DOCKER_PROD) login --password $(GITLAB_ACCESS_TOKEN) --username $(GITLAB_USERNAME) $(REGISTRY)
	$(DC_PROD) pull
	$(DC_PROD) up --detach
prod.down:
	$(DC_PROD) down
prod.logs:
	$(DC_PROD) logs --follow --tail 1000
prod.ssh:
	ssh -K $(PROD_HOST)
prod.network.create:
	$(DOCKER_PROD) network create $(NETWORK)
prod.network.remove:
	$(DOCKER_PROD) network remove $(NETWORK)
