package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/evermos/boilerplate-go/internal/domain/products"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	ProductService products.ProductService
}

func ProvideProductHandler(Productervice products.ProductService) ProductHandler {
	return ProductHandler{
		ProductService: Productervice,
	}
}

func (h *ProductHandler) Router(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Post("/", h.CreateProduct)
	})
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestFormat products.PayloadProductAndVariant
	err := decoder.Decode(&requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	err = shared.GetValidator().Struct(requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	prod, err := h.ProductService.CreateWithVariant(requestFormat)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, prod)
}
