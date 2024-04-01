#! /usr/bin/make

MODULE=$(shell go list -m)
MIGRATION_PATH = $(CURDIR)/pkg/databasefx/migrations
POSTGRESQL_URL = postgres://postgres:postgres@localhost:5432/advertise?sslmode=disable

export POSTGRESQL_URL

gen:
	@echo GENERATING CODE...
	@goa gen $(MODULE)/api/design -o api

migrate:
	@migrate -database ${POSTGRESQL_URL} -path $(MIGRATION_PATH) up

create_migration cm:
ifdef NAME
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(NAME)
else
	@echo Please provide a migration name using 'make create_migration NAME=your_migration_name'
endif