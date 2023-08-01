package variants

import (
	"encoding/json"
	"time"

	"github.com/evermos/boilerplate-go/internal/domain/image"
	"github.com/evermos/boilerplate-go/shared"
	"github.com/evermos/boilerplate-go/shared/nuuid"
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
	Price       float64       `db:"price"`
	Status      string        `db:"status"`
	Quantity    int           `db:"quantity"`
	Images      []image.Image `db:"-"`
	CreatedAt   time.Time     `db:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at"`
	DeletedAt   null.Time     `db:"deleted_at"`
	CreatedBy   uuid.UUID     `db:"created_by"`
	UpdatedBy   uuid.UUID     `db:"updated_by"`
	DeletedBy   nuuid.NUUID   `db:"deleted_by"`
}

type PayloadVariant struct {
	VariantName  string   `json:"variantName"`
	Price        float64  `json:"price"`
	Status       string   `json:"status"`
	Quantity     int      `json:"quantity"`
	ImagePayload []string `json:"images"`
}

type VariantResponseFormat struct {
	VariantID   uuid.UUID   `json:"variantId"`
	ProductID   uuid.UUID   `json:"productId"`
	VariantName string      `json:"variantName"`
	Price       float64     `json:"price"`
	Status      string      `json:"status"`
	Quantity    int         `json:"quantity"`
	Images      []string    `json:"images"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	DeletedAt   null.Time   `json:"deletedAt"`
	CreatedBy   uuid.UUID   `json:"createdBy"`
	UpdatedBy   uuid.UUID   `json:"updatedBy"`
	DeletedBy   nuuid.NUUID `json:"deletedBy"`
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

	var urlsOnly []string
	for i := range v.Images {
		urlsOnly = append(urlsOnly, v.Images[i].ImageURL)
	}
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
		Images:      urlsOnly,
	}
	return resp
}

func (v Variant) NewFromPayload(payload PayloadVariant, proId uuid.UUID) (Variant, error) {
	var img image.Image
	varId, _ := uuid.NewV4()
	var imgs []image.Image
	for _, link := range payload.ImagePayload {
		pay, err := img.NewFromPayload(link, varId)
		if err != nil {
			return Variant{}, err
		}
		imgs = append(imgs, pay)
	}
	newVar := Variant{
		ProductID:   proId,
		VariantID:   varId,
		VariantName: payload.VariantName,
		Price:       payload.Price,
		Images:      imgs,
		Status:      payload.Status,
		Quantity:    payload.Quantity,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   proId,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   proId,
	}

	err := newVar.Validate()
	return newVar, err
}

func (p *Variant) Validate() error {
	validator := shared.GetValidator()
	return validator.Struct(p)
}

func (v Variant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.ToResponseFormat())
}
