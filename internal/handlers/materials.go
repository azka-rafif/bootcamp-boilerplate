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
	})
}

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
