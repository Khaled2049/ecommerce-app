package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Khaled2049/ecommerce-app/internal/handler"
	"github.com/Khaled2049/ecommerce-app/internal/repository/postgres"
	"github.com/Khaled2049/ecommerce-app/internal/service"
	"github.com/Khaled2049/ecommerce-app/pkg/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Initialize database
	dbConfig := database.NewConfig(dbURL)
	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatal("Error connecting to db", err)
	}

	defer db.Close()
	log.Println("✅ Successfully connected to database!")

	// Initialize repositories
	customerRepo := postgres.NewCustomerRepository(db)
	orderRepo := postgres.NewOrderRepository(db)

	// Initialize services
	customerService := service.NewCustomerService(customerRepo)
	orderService := service.NewOrderService(orderRepo)

	// Initialize handlers
	customerHandler := handler.NewCustomerHandler(customerService)
	orderHandler := handler.NewOrderHandler(orderService)

	// Router setup
	router := mux.NewRouter()

	// Register routes
	customerHandler.RegisterRoutes(router)
	orderHandler.RegisterRoutes(router)

	// Add middleware
	router.Use(loggingMiddleware)

	// Server configuration
	port := ":8080"

	// Add informative startup message
	log.Printf("Server starting on http://localhost%s 🚀", port)
	log.Printf("  Available endpoints:")
	log.Printf("  POST   /customers")
	log.Printf("  GET    /customers/{id}")
	log.Printf("  GET    /customers")
	log.Printf("  PUT    /customers/{id}")
	log.Printf("  DELETE /customers/{id}")
	// Order routes
	log.Printf("  POST   /orders")
	log.Printf("  GET    /orders/{id}")
	log.Printf("  GET    /orders")
	log.Printf("  PUT    /orders/{id}")
	log.Printf("  DELETE /orders/{id}")

	// Start server
	log.Fatal(http.ListenAndServe(port, router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
