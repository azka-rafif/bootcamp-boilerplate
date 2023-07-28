package brand

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type Brand struct {
	BrandID   uuid.UUID
	BrandName string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	DeletedBy uuid.UUID
}
