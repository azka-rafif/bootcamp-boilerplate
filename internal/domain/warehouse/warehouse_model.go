package warehouse

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type Warehouse struct {
	WarehouseID   uuid.UUID
	WarehouseName string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     null.Time
	CreatedBy     uuid.UUID
	UpdatedBy     uuid.UUID
	DeletedBy     uuid.UUID
}
