#!/bin/bash

# Base directory
# mkdir -p ecommerce

# cmd directory
mkdir -p cmd/api
touch cmd/api/main.go

# internal directory
mkdir -p internal/config
touch internal/config/config.go

mkdir -p internal/domain
touch internal/domain/customer.go
touch internal/domain/order.go
touch internal/domain/product.go
touch internal/domain/payment.go

mkdir -p internal/repository/postgres
touch internal/repository/postgres/customer.go
touch internal/repository/postgres/order.go
touch internal/repository/postgres/product.go
touch internal/repository/postgres/payment.go

mkdir -p internal/repository/interfaces
touch internal/repository/interfaces/customer.go
touch internal/repository/interfaces/order.go
touch internal/repository/interfaces/product.go
touch internal/repository/interfaces/payment.go

mkdir -p internal/service
touch internal/service/customer.go
touch internal/service/order.go
touch internal/service/product.go
touch internal/service/payment.go

mkdir -p internal/handler
touch internal/handler/customer.go
touch internal/handler/order.go
touch internal/handler/product.go
touch internal/handler/payment.go

# pkg directory
mkdir -p pkg/logger
touch pkg/logger/logger.go

mkdir -p pkg/middleware
touch pkg/middleware/auth.go
touch pkg/middleware/logging.go

# migrations directory
mkdir -p migrations
touch migrations/schema.sql

# Go module files
touch go.mod
touch go.sum

echo "Project structure created successfully!"
