package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/service"
	"github.com/Khaled2049/ecommerce-app/pkg/utils"
	"github.com/gorilla/mux"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/orders", h.Create).Methods("POST")
	r.HandleFunc("/orders/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/orders", h.List).Methods("GET")
	r.HandleFunc("/orders/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/orders/{id}", h.Delete).Methods("DELETE")
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var orderCreate domain.OrderCreate
	if err := json.NewDecoder(r.Body).Decode(&orderCreate); err != nil {
		utils.WriteError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	created, err := h.orderService.CreateOrder(r.Context(), &orderCreate)
	if err != nil {
		utils.WriteError(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, created)
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.WriteError(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.orderService.GetOrderByID(r.Context(), orderID)
	if err != nil {
		utils.WriteError(w, "Order not found", http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.GetOrders(r.Context(), 10, 0)
	if err != nil {
		utils.WriteError(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orders)
}

func (h *OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.WriteError(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.WriteError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	order.OrderID = orderID

	updated, err := h.orderService.UpdateOrder(r.Context(), &order)
	if err != nil {
		utils.WriteError(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, updated)
}

func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.WriteError(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	err = h.orderService.DeleteOrder(r.Context(), orderID)
	if err != nil {
		utils.WriteError(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
