# Makefile for todo-be project

# Variables
include .env
migrate-up:
	@echo "Running migrations (up)..."
	@goose -dir ./db/migrations postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} sslmode=disable" up

# Revert database migrations (down)
migrate-down:
	@echo "Reverting migrations (down)..."
	@goose -dir ./db/migrations postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable" down

.PHONY: migrate-up migrate-down