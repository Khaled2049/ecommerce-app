package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/service"
	"github.com/gorilla/mux"
)

type PaymentHandler struct {
	paymentService *service.PaymentService
}

func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

func (h *PaymentHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/payments", h.Create).Methods("POST")
	r.HandleFunc("/payments/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/payments", h.List).Methods("GET")
	r.HandleFunc("/payments/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/payments/{id}", h.Delete).Methods("DELETE")
}

func (h *PaymentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.PaymentCreate

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment, err := h.paymentService.CreatePayment(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func (h *PaymentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	paymentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	payment, err := h.paymentService.GetPaymentByID(r.Context(), paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func (h *PaymentHandler) List(w http.ResponseWriter, r *http.Request) {
	payments, err := h.paymentService.GetPayments(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

func (h *PaymentHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	paymentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	var input domain.PaymentUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.ID = paymentID
	payment, err := h.paymentService.UpdatePayment(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func (h *PaymentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	paymentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	if err := h.paymentService.DeletePayment(r.Context(), paymentID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
