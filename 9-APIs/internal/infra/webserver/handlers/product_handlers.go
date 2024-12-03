package handlers

import (
	"encoding/json"
	"github.com/eduardo1520/goexpert/9-APIs/internal/dto"
	"github.com/eduardo1520/goexpert/9-APIs/internal/entity"
	"github.com/eduardo1520/goexpert/9-APIs/internal/infra/database"
	entityPkg "github.com/eduardo1520/goexpert/9-APIs/pkg/entity"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ErrorProduct struct {
	Message string `json:"message"`
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct godoc
// @Summary Cria um novo produto
// @Description Adiciona um produto ao banco de dados com base nos dados enviados na requisição
// @Tags products
// @Accept json
// @Produce json
// @Param request body dto.CreateProductInput true "product request"
// @Success 201
// @Failure 400 {object} ErrorProduct
// @Failure 500 {object} ErrorProduct
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, float64(product.Price))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary Get product by ID
// @Description Retrieve a product using its unique ID.
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 404 {object} ErrorProduct
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// GetProducts godoc
// @Summary Get all products
// @Description Retrieve a list of products with pagination and sorting options.
// @Tags products
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "limit"
// @Success 200 {array} entity.Product
// @Failure 404 {object} ErrorProduct
// @Failure 500 {object} ErrorProduct
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product's details by its ID.
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" Format(uuid)
// @Param request body dto.CreateProductInput true "Product request"
// @Success 200
// @Failure 404
// @Failure 500 {object} ErrorProduct
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Remove an existing product by its ID.
// @Tags products
// @Param id path string true "Product ID" Format(uuid)
// @Success 200
// @Failure 404
// @Failure 500 {object} ErrorProduct
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
