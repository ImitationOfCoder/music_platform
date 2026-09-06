include .env
export

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

clear-logs:
	@rm ./services/*/out/logs/*.log

dev:
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml up

build-dev:
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml build

prod:
	@docker compose -f docker-compose.yml up

build-prod:
	@docker compose -f docker-compose.yml build
