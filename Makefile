-include .env
export

DOCKER_COMPOSE_FILE=docker-compose.yaml


# go generate	
.PHONY: go-gen
go-gen:
	go generate ./...

.PHONY: start
start:
	@echo "Start Containers"
	sudo docker compose -f ${DOCKER_COMPOSE_FILE} up -d ${DOCKER_SERVICES}
	sleep 2
	sudo docker compose -f ${DOCKER_COMPOSE_FILE} ps

.PHONY: stop
stop:
	@echo "Stop Containers"
	sudo docker compose -f ${DOCKER_COMPOSE_FILE} stop ${DOCKER_SERVICES}
	sleep 2
	sudo docker compose -f ${DOCKER_COMPOSE_FILE} ps

.PHONY: stop
rm: stop
	@echo "Remove Containers"
	sudo docker compose -f ${DOCKER_COMPOSE_FILE} rm -v -f ${DOCKER_SERVICES}

.PHONY: migration-up
migration-up:
	@echo "Migrations Up"
	sleep 2
	sudo docker compose run --rm migrate -path=migrations/ -database='postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable'  force 2
	sudo docker compose run --rm migrate -path=migrations/ -database='postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' up

.PHONY: migration-generate
migration-generate:
	@echo "Generation migration file $(name)"
	sleep 2
	sudo docker compose run --rm migrate create -ext sql -dir ./migrations -seq $(name)