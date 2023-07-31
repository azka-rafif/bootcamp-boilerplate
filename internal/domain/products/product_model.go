package products

import (
	"encoding/json"
	"time"

	"github.com/evermos/boilerplate-go/internal/domain/variants"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/nuuid"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type Product struct {
	ProductID   uuid.UUID   `db:"product_id" validate:"required"`
	UserID      uuid.UUID   `db:"user_id" validate:"required"`
	BrandID     uuid.UUID   `db:"brand_id" validate:"required"`
	ProductName string      `db:"product_name" validate:"required"`
	CreatedAt   time.Time   `db:"created_at" validate:"required"`
	UpdatedAt   time.Time   `db:"updated_at" validate:"required"`
	DeletedAt   null.Time   `db:"deleted_at"`
	CreatedBy   uuid.UUID   `db:"created_by" validate:"required"`
	UpdatedBy   uuid.UUID   `db:"updated_by" validate:"required"`
	DeletedBy   nuuid.NUUID `db:"deleted_by"`
}

type PayloadProductAndVariant struct {
	UserID         uuid.UUID               `json:"userId" validate:"required"`
	BrandID        uuid.UUID               `json:"brandId" validate:"required"`
	ProductName    string                  `json:"productName" validate:"required"`
	VariantPayload variants.PayloadVariant `json:"variant" validate:"required"`
}

type PayloadProduct struct {
	UserID      uuid.UUID `json:"userId" validate:"required"`
	BrandID     uuid.UUID `json:"brandId" validate:"required"`
	ProductName string    `json:"productName" validate:"required"`
}

type ProductResponseFormat struct {
	ProductID   uuid.UUID   `json:"productId"`
	UserID      uuid.UUID   `json:"userId"`
	BrandID     uuid.UUID   `json:"brandId"`
	ProductName string      `json:"productName"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	DeletedAt   null.Time   `json:"deletedAt"`
	CreatedBy   uuid.UUID   `json:"createdBy"`
	UpdatedBy   uuid.UUID   `json:"updatedBy"`
	DeletedBy   nuuid.NUUID `json:"deletedBy"`
}

type ProductAndVariant struct {
	Product Product
	Variant variants.Variant
}

type ProductAndVariantResponseFormat struct {
	Product ProductResponseFormat
	Variant variants.VariantResponseFormat
}

type ProductWithVariants struct {
	Product  Product
	Variants []variants.Variant
}

type ProductWithVariantsResponseFormat struct {
	Product  Product
	Variants []variants.Variant `json:"variants"`
}

func (pv ProductAndVariant) NewFromPayload(payload PayloadProductAndVariant) (ProductAndVariant, error) {
	proId, _ := uuid.NewV4()
	newPro := Product{
		ProductID:   proId,
		UserID:      payload.UserID,
		BrandID:     payload.BrandID,
		ProductName: payload.ProductName,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   payload.UserID,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   payload.UserID,
	}
	varId, _ := uuid.NewV4()
	newVar := variants.Variant{
		VariantID:   varId,
		ProductID:   proId,
		VariantName: payload.VariantPayload.VariantName,
		Price:       payload.VariantPayload.Price,
		Status:      "ready",
		Quantity:    payload.VariantPayload.Quantity,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   payload.UserID,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   payload.UserID,
	}
	res := ProductAndVariant{Product: newPro, Variant: newVar}
	err := newPro.Validate()
	return res, err
}

func (p Product) NewFromPayload(payload PayloadProduct) (Product, error) {
	proId, _ := uuid.NewV4()
	newPro := Product{
		ProductID:   proId,
		UserID:      payload.UserID,
		BrandID:     payload.BrandID,
		ProductName: payload.ProductName,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   payload.UserID,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   payload.UserID,
	}

	err := newPro.Validate()
	return newPro, err
}

func (pv *ProductWithVariants) ToResponseFormat() ProductWithVariantsResponseFormat {

	for i, varis := range pv.Variants {
		pv.Variants[i] = variants.Variant(varis.ToResponseFormat())
	}

	resp := ProductWithVariantsResponseFormat{
		Product:  Product(pv.Product),
		Variants: pv.Variants,
	}
	return resp
}

func (pv *ProductAndVariant) ToResponseFormat() ProductAndVariantResponseFormat {
	resp := ProductAndVariantResponseFormat{
		Product: ProductResponseFormat(pv.Product),
		Variant: pv.Variant.ToResponseFormat(),
	}
	return resp
}

func (p *Product) ToResponseFormat() ProductResponseFormat {
	resp := ProductResponseFormat{
		ProductID:   p.ProductID,
		UserID:      p.UserID,
		BrandID:     p.BrandID,
		ProductName: p.ProductName,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   p.UserID,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   p.UserID,
	}
	return resp
}

func (p *Product) Validate() error {
	validator := shared.GetValidator()
	return validator.Struct(p)
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.ToResponseFormat())
}

func (p *Product) Update(req PayloadProduct) (err error) {
	p.UpdatedAt = time.Now().UTC()
	p.UpdatedBy = req.UserID
	p.ProductName = req.ProductName
	p.BrandID = req.BrandID

	err = p.Validate()

	return
}
