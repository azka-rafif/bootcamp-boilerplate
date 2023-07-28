package products

import "github.com/evermos/boilerplate-go/configs"

type ProductService interface {
	CreateWithVariant(newMat PayloadProductAndVariant) (ProductAndVariant, error)
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
