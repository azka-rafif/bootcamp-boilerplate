package materials

import (
	"github.com/evermos/boilerplate-go/configs"
	"github.com/evermos/boilerplate-go/event/producer"
)

type MaterialService interface {
	Create(newMat PayloadMaterial) (Material, error)
	GetAll(limit, page int, sort, field string) ([]Material, error)
}

type MaterialServiceImpl struct {
	MaterialRepository MaterialRepository
	Producer           producer.Producer
	Config             *configs.Config
}

func ProvideMaterialServiceImpl(materialRepo MaterialRepository, producer producer.Producer, config *configs.Config) *MaterialServiceImpl {
	s := new(MaterialServiceImpl)
	s.MaterialRepository = materialRepo
	s.Config = config
	s.Producer = producer

	return s
}

func (s *MaterialServiceImpl) Create(payload PayloadMaterial) (Material Material, err error) {

	err = s.MaterialRepository.Create(payload)
	if err != nil {
		return
	}
	return
}
