package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/service"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(customerService *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/customers", h.Create).Methods("POST")
	// r.HandleFunc("/customers/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/customers", h.List).Methods("GET")
	// r.HandleFunc("/customers/{id}", h.Update).Methods("PUT")
	// r.HandleFunc("/customers/{id}", h.Delete).Methods("DELETE")
}

func (h *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CustomerCreate

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := h.customerService.CreateCustomer(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerHandler) List(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.GetCustomers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
