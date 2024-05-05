package catusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sCatUsecase) Update(catId *int, p *entities.ParamsUpdateCat) (*ResultCreate, error) {
	filters := entities.CatSearchFilter{
		ID: *catId,
	}

	cat, _ := i.catRepository.IsExists(&filters)

	if !cat {
		return nil, ErrCatNotFound
	}

	data, err := i.catRepository.Update(catId,
		&entities.ParamsUpdateCat{
			Name:        p.Name,
			Race:        p.Race,
			Sex:         p.Sex,
			AgeInMonth:  p.AgeInMonth,
			Description: p.Description,
			ImageUrls:   p.ImageUrls,
		},
	)

	if err != nil {
		return nil, err
	}

	return &ResultCreate{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
	}, nil

}
