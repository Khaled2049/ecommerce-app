.PHONY: db-init db-seed db-reset

db-init:
	docker exec -i postgres_bootcamp psql -U bootcamp_user -d ecommerce_db < migrations/schema.sql

db-seed:
	docker exec -i postgres_bootcamp psql -U bootcamp_user -d ecommerce_db < migrations/seed.sql

db-reset:
	docker-compose down -v
	docker-compose up -d