package image

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type Image struct {
	ImageID   uuid.UUID
	VariantID uuid.UUID
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	DeletedBy uuid.UUID
}
