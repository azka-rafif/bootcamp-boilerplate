package materials

import (
	"github.com/evermos/boilerplate-go/configs"
)

type MaterialService interface {
	Create(newMat PayloadMaterial) (Material, error)
	GetAll() (mats []Material, err error)
}

type MaterialServiceImpl struct {
	MaterialRepository MaterialRepository
	Config             *configs.Config
}

func ProvideMaterialServiceImpl(materialRepo MaterialRepository, config *configs.Config) *MaterialServiceImpl {
	s := new(MaterialServiceImpl)
	s.MaterialRepository = materialRepo
	s.Config = config

	return s
}

func (s *MaterialServiceImpl) Create(payload PayloadMaterial) (material Material, err error) {
	material, err = material.NewFromPayload(payload)
	if err != nil {
		return
	}
	err = s.MaterialRepository.Create(material)
	if err != nil {
		return
	}
	return
}

func (s *MaterialServiceImpl) GetAll() (mats []Material, err error) {
	mats, err = s.MaterialRepository.GetAll()
	if err != nil {
		return
	}
	return
}
