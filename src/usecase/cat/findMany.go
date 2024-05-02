package catusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sCatUsecase) FindMany(filters *entities.CatSearchFilter) ([]*entities.Cat, error) {
	allCats, err := i.catRepository.FindMany(filters)

	if err != nil {
		return nil, err
	}

	return allCats, nil
}
