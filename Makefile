MIGRATIONS_DIR := migrations
DB_FILE := pasta.sqlite
DB_DRIVER := sqlite3
GOOSE := goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_FILE)

.PHONY: run create up down status

run:
	@go run .

create:
ifndef NAME
	$(error NAME is not set. Usage: make create NAME=<migration_name>)
endif
	$(GOOSE) create $(NAME) sql

up:
	$(GOOSE) up

down:
	$(GOOSE) down

status:
	$(GOOSE) status
