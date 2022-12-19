include .env
export

NAME := notion_service
NETWORK := notion_service
DOCKER := docker
DC := docker-compose

# REGISTRY=gitlab-registry.42paris.fr

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

validate:
	$(DOCKER) run --rm \
		--volume $(shell pwd)/change_checker_go:/go/src/gitlab.42paris.fr/notion_service \
		--workdir /go/src/gitlab.42paris.fr/notion_service \
		quay.io/goswagger/swagger validate ./gen/swagger.yml

gen:
	rm -frv ./change_checker_go/pkg/models ./change_checker_go/pkg/api/operations
	$(DOCKER) run --rm \
		--volume $(shell pwd)/change_checker_go:/go/src/gitlab.42paris.fr/notion_service \
		--workdir /go/src/gitlab.42paris.fr/notion_service \
		quay.io/goswagger/swagger generate server --name $(NAME) --spec ./gen/swagger.yml --target ./pkg --main-package=../../cmd/notion-server/
