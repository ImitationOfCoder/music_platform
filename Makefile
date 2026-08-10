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
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disabel \
		"$(action)"
