package image

import (
	"time"

	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/nuuid"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type Image struct {
	ImageID   uuid.UUID   `db:"image_id"`
	VariantID uuid.UUID   `db:"variant_id"`
	ImageURL  string      `db:"image_url"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
	DeletedAt null.Time   `db:"deleted_at"`
	CreatedBy uuid.UUID   `db:"created_by"`
	UpdatedBy uuid.UUID   `db:"updated_by"`
	DeletedBy nuuid.NUUID `db:"deleted_by"`
}

func (i Image) NewFromPayload(urls string, variantId uuid.UUID) (Image, error) {
	imgId, _ := uuid.NewV4()
	newImg := Image{
		ImageID:   imgId,
		VariantID: variantId,
		ImageURL:  urls,
		CreatedAt: time.Now().UTC(),
		CreatedBy: variantId,
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: variantId,
	}
	err := newImg.Validate()

	return newImg, err
}

func (i *Image) Validate() error {
	validator := shared.GetValidator()
	return validator.Struct(i)
}
