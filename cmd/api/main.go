package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Khaled2049/ecommerce-app/internal/handler"
	"github.com/Khaled2049/ecommerce-app/internal/repository/postgres"
	"github.com/Khaled2049/ecommerce-app/internal/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  Warning: .env file not found")
	}

	// Get database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ DATABASE_URL environment variable is not set")
	}

	// Database connection
	log.Printf("📡 Attempting to connect to database...")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("❌ Failed to open database connection: ", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatal("❌ Could not connect to database: ", err)
	}
	log.Println("✅ Successfully connected to database!")
	// Initialize repositories
	customerRepo := postgres.NewCustomerRepository(db)

	// Initialize services
	customerService := service.NewCustomerService(customerRepo)

	// Initialize handlers
	customerHandler := handler.NewCustomerHandler(customerService)

	// Router setup
	router := mux.NewRouter()

	// Register routes
	customerHandler.RegisterRoutes(router)

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

	// Start server
	log.Fatal(http.ListenAndServe(port, router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
