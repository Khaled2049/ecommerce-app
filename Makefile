.PHONY: db-init db-seed db-reset

db-init:
	docker exec -i ecommerce-postgres psql -U bootcamp_user -d ecommerce_db < migrations/01-schema.sql

db-seed:
	docker exec -i ecommerce-postgres psql -U bootcamp_user -d ecommerce_db < migrations/02-seed.sql

db-reset:
	docker-compose down -v
	docker-compose up -d