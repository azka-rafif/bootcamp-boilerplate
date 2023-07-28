package variants

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type VariantStatus int

const (
	Ready VariantStatus = iota
	OutOfStock
	Limited
)

type Variant struct {
	VariantID   uuid.UUID     `db:"variant_id"`
	ProductID   uuid.UUID     `db:"product_id"`
	VariantName string        `db:"variant_name"`
	Price       int           `db:"price"`
	Status      VariantStatus `db:"status"`
	Quantity    int           `db:"quantity"`
	CreatedAt   time.Time     `db:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at"`
	DeletedAt   null.Time     `db:"deleted_at"`
	CreatedBy   uuid.UUID     `db:"created_by"`
	UpdatedBy   uuid.UUID     `db:"updated_by"`
	DeletedBy   uuid.UUID     `db:"deleted_by"`
}

type PayloadVariant struct {
	VariantName string `json:"variantName"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
	Quantity    int    `json:"quantity"`
}

type VariantResponseFormat struct {
	VariantID   uuid.UUID     `json:"variantId"`
	ProductID   uuid.UUID     `json:"productId"`
	VariantName string        `json:"variantName"`
	Price       int           `json:"price"`
	Status      VariantStatus `json:"status"`
	Quantity    int           `json:"quantity"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	DeletedAt   null.Time     `json:"deletedAt"`
	CreatedBy   uuid.UUID     `json:"createdBy"`
	UpdatedBy   uuid.UUID     `json:"updatedBy"`
	DeletedBy   uuid.UUID     `json:"deletedBy"`
}

func GetVariantStatus(stat VariantStatus) string {

	switch stat {
	case Ready:
		return "ready"
	case OutOfStock:
		return "out_of_stock"
	case Limited:
		return "limited"
	default:
		return "ready"
	}
}

func (v *Variant) ToResponseFormat() VariantResponseFormat {
	resp := VariantResponseFormat{
		VariantID:   v.VariantID,
		ProductID:   v.ProductID,
		VariantName: v.VariantName,
		Price:       v.Price,
		Quantity:    v.Quantity,
		Status:      v.Status,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
		CreatedBy:   v.CreatedBy,
		UpdatedBy:   v.UpdatedBy,
		DeletedBy:   v.DeletedBy,
	}
	return resp
}
