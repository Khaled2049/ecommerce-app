.PHONY: db-reset

db-reset:
	docker-compose down -v
	docker-compose up -d