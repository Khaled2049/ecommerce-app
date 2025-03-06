package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Khaled2049/ecommerce-app/internal/domain"
	"github.com/Khaled2049/ecommerce-app/internal/service"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", h.Create).Methods("POST")
	r.HandleFunc("/products/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/products", h.List).Methods("GET")
	r.HandleFunc("/products/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/products/{id}", h.Delete).Methods("DELETE")
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.ProductCreate

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.productService.CreateProduct(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.productService.GetProductByID(r.Context(), productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.ListProducts(r.Context(), 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var input domain.ProductUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.ID = productID
	product, err := h.productService.UpdateProduct(r.Context(), &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := h.productService.DeleteProduct(r.Context(), productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
