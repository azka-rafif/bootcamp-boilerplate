package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/evermos/boilerplate-go/internal/domain/products"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/shared/pagination"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
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
		r.Get("/", h.GetAllProducts)
		r.Post("/", h.CreateProduct)
		r.Post("/add-variant/{id}", h.AddVariants)
		r.Get("/{id}", h.GetProductByID)
		r.Put("/{id}", h.UpdateProduct)
		r.Delete("/soft/{id}", h.SoftDelete)
		r.Delete("/hard/{id}", h.HardDelete)
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

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page, err := pagination.ConvertToInt(pagination.ParseQueryParams(r, "page"))
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	limit, err := pagination.ConvertToInt(pagination.ParseQueryParams(r, "limit"))
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	sort := pagination.GetSortDirection(pagination.ParseQueryParams(r, "sort"))
	field := pagination.CheckFieldQuery(pagination.ParseQueryParams(r, "field"), "product_id")

	pg := pagination.NewPaginationQuery(page, limit, field, sort)

	prods, err := h.ProductService.GetAllProducts(pg)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, prods)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := uuid.FromString(idString)

	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	prod, err := h.ProductService.GetProductByID(id)

	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, prod)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := uuid.FromString(idString)

	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestFormat products.PayloadProduct
	err = decoder.Decode(&requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	err = shared.GetValidator().Struct(requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	prod, err := h.ProductService.Update(id, requestFormat)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, prod)
}

func (h *ProductHandler) SoftDelete(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := uuid.FromString(idString)

	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	decoder := json.NewDecoder(r.Body)
	var requestFormat products.PayloadProduct
	err = decoder.Decode(&requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	err = shared.GetValidator().Struct(requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	prod, err := h.ProductService.SoftDelete(id, requestFormat)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, prod)
}

func (h *ProductHandler) HardDelete(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := uuid.FromString(idString)

	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	decoder := json.NewDecoder(r.Body)
	var requestFormat products.PayloadProduct
	err = decoder.Decode(&requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	err = shared.GetValidator().Struct(requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	err = h.ProductService.HardDelete(id, requestFormat)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusNoContent, nil)
}

func (h *ProductHandler) AddVariants(w http.ResponseWriter, r *http.Request) {

}
