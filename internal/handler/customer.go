package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/service"
	"github.com/Khaled2049/ecommerce-app/pkg/utils"
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
	r.HandleFunc("/customers/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/customers", h.List).Methods("GET")
	r.HandleFunc("/customers/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/customers/{id}", h.Delete).Methods("DELETE")
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

func (h *CustomerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customerID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	customer, err := h.customerService.GetCustomerByID(r.Context(), customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customerID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.WriteError(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	var input domain.CustomerUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the ID from the URL parameter
	input.ID = customerID

	// Update the customer
	updated, err := h.customerService.UpdateCustomer(r.Context(), &input)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrCustomerNotFound):
			utils.WriteError(w, "Customer not found", http.StatusNotFound)
		case errors.Is(err, domain.ErrEmailAlreadyExists):
			utils.WriteError(w, "Email already exists", http.StatusConflict)
		default:
			utils.WriteError(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, updated)
}

func (h *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customerID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.WriteError(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	err = h.customerService.DeleteCustomer(r.Context(), customerID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrCustomerNotFound):
			utils.WriteError(w, "Customer not found", http.StatusNotFound)
		default:
			utils.WriteError(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
