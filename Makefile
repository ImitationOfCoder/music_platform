include .env
export

up:
	docker compose up --watch

down:
	docker compose down

build:
	docker compose build

create-migration:
	@if [ -z "$(seq)" ]; then \
		echo "Variable 'seq' is empty."; \
		exit 1; \
	fi; \
	docker compose run --rm postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	make migrate-action action=up

migrate-down:
	make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Variable 'action' is empty."; \
		exit 1; \
	fi; \
	docker compose run --rm postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable \
		"$(action)"

run:
	@export LOGGER_FOLDER=./out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run -mod=mod cmd/music_platform/main.go

clear-logs:
	@rm ./out/logs/*.log

dev:
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml up

build-dev:
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml build

prod:
	@docker compose -f docker-compose.yml up

build-prod:
	@docker compose -f docker-compose.yml build
