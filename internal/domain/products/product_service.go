package products

import (
	"github.com/evermos/boilerplate-go/configs"
	"github.com/evermos/boilerplate-go/internal/domain/variants"
	"github.com/evermos/boilerplate-go/shared/pagination"
	"github.com/gofrs/uuid"
)

type ProductService interface {
	CreateWithVariant(newMat PayloadProductAndVariant) (ProductAndVariant, error)
	GetAllProducts(pg pagination.Pagination) (prods []Product, err error)
	GetProductByID(prodId uuid.UUID) (prod ProductWithVariants, err error)
	Update(prodId uuid.UUID, payload PayloadProduct) (prod Product, err error)
	SoftDelete(prodId uuid.UUID, payload PayloadProduct) (prod Product, err error)
	HardDelete(prodId uuid.UUID, payload PayloadProduct) (err error)
	AddVariant(prodId uuid.UUID, payload variants.PayloadVariant) (variant variants.Variant, err error)
}

type ProductServiceImpl struct {
	ProductRepository ProductRepository
	Config            *configs.Config
}

func ProvideProductServiceImpl(ProductRepo ProductRepository, config *configs.Config) *ProductServiceImpl {
	s := new(ProductServiceImpl)
	s.ProductRepository = ProductRepo
	s.Config = config

	return s
}

func (s *ProductServiceImpl) CreateWithVariant(payload PayloadProductAndVariant) (ProductAndVariant ProductAndVariant, err error) {
	ProductAndVariant, err = ProductAndVariant.NewFromPayload(payload)
	if err != nil {
		return
	}
	err = s.ProductRepository.CreateWithVariant(ProductAndVariant)
	if err != nil {
		return
	}
	return
}

func (s *ProductServiceImpl) GetAllProducts(pg pagination.Pagination) (prods []Product, err error) {
	prods, err = s.ProductRepository.GetAllProducts(pg.Field, pg.Sort, pg.Limit, pg.Offset)

	if err != nil {
		return
	}

	return
}

func (s *ProductServiceImpl) GetProductByID(prodId uuid.UUID) (prod ProductWithVariants, err error) {
	prod, err = s.ProductRepository.GetProductWithVariants(prodId)
	if err != nil {
		return
	}
	println(prod.Variants[0].Images)
	return
}

func (s *ProductServiceImpl) Update(prodId uuid.UUID, payload PayloadProduct) (prod Product, err error) {
	prod, err = s.ProductRepository.GetProductByID(prodId)
	if err != nil {
		return
	}
	err = prod.Update(payload)
	if err != nil {
		return
	}
	err = s.ProductRepository.Update(prod)
	return
}

func (s *ProductServiceImpl) SoftDelete(prodId uuid.UUID, payload PayloadProduct) (prod Product, err error) {
	prod, err = s.ProductRepository.GetProductByID(prodId)
	if err != nil {
		return
	}

	err = prod.SoftDelete(payload.UserID)
	if err != nil {
		return
	}
	err = s.ProductRepository.Update(prod)
	return
}

func (s *ProductServiceImpl) HardDelete(prodId uuid.UUID, payload PayloadProduct) (err error) {
	_, err = s.ProductRepository.GetProductByID(prodId)
	if err != nil {
		return
	}
	err = s.ProductRepository.HardDelete(prodId)
	return
}

func (s *ProductServiceImpl) AddVariant(prodId uuid.UUID, payload variants.PayloadVariant) (variant variants.Variant, err error) {
	variant, err = variant.NewFromPayload(payload, prodId)
	if err != nil {
		return
	}
	err = s.ProductRepository.AddVariant(variant)
	return
}
