package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/evermos/boilerplate-go/internal/domain/materials"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
)

type MaterialsHandler struct {
	MaterialService materials.MaterialService
}

func ProvideMaterialsHandler(materialService materials.MaterialService) MaterialsHandler {
	return MaterialsHandler{
		MaterialService: materialService,
	}
}

func (h *MaterialsHandler) Router(r chi.Router) {
	r.Route("/materials", func(r chi.Router) {
		r.Post("/", h.CreateMaterial)
		r.Get("/", h.GetAllMaterial)
	})
}

// CreateMaterial creates a new Material.
// @Summary Create a new Material.
// @Description This endpoint creates a new Material.
// @Tags materials/material
// @Param foo body materials.PayloadMaterial true "The Material to be created."
// @Produce json
// @Success 201 {object} response.Base{data=materials.MaterialResponseFormat}
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/materials [post]
func (h *MaterialsHandler) CreateMaterial(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestFormat materials.PayloadMaterial
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
	mat, err := h.MaterialService.Create(requestFormat)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, mat)
}

func (h *MaterialsHandler) GetAllMaterial(w http.ResponseWriter, r *http.Request) {
	mats, err := h.MaterialService.GetAll()
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, mats)
}
