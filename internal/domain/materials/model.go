package materials

import (
	"time"

	"github.com/evermos/boilerplate-go/shared"
	"github.com/gofrs/uuid"
)

type Material struct {
	Id          uuid.UUID `db:"id" validate:"required"`
	Title       string    `db:"title" validate:"required"`
	Description string    `db:"description" validate:"required"`
	CreatedAt   time.Time `db:"created_at" validate:"required"`
	UpdatedAt   time.Time `db:"updated_at" validate:"required"`
}

type PayloadMaterial struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type MaterialResponseFormat struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (m Material) NewFromPayload(payload PayloadMaterial) (newMat Material, err error) {
	matID, _ := uuid.NewV4()
	newMat = Material{
		Id:          matID,
		Title:       payload.Title,
		Description: payload.Description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	err = m.Validate()
	return
}

func (m Material) ToResponseFormat() MaterialResponseFormat {
	resp := MaterialResponseFormat{
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
	return resp
}

func (m *Material) Validate() error {
	validator := shared.GetValidator()
	return validator.Struct(m)
}
