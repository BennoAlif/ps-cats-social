package catusecase

import (
	"time"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	userusecase "github.com/BennoAlif/ps-cats-social/src/usecase/user"
)

type (
	ResultCreate struct {
		ID        int64     `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
	}
)

func (i *sCatUsecase) Create(p *entities.ParamsCreateCat) (*ResultCreate, error) {
	filters := entities.ParamsCreateUser{
		ID: int64(p.UserId),
	}

	user, _ := i.userRepository.IsExists(&filters)

	if !user {
		return nil, userusecase.ErrInvalidUser
	}

	data, err := i.catRepository.Create(&entities.ParamsCreateCat{
		Name:        p.Name,
		Race:        p.Race,
		Sex:         p.Sex,
		AgeInMonth:  p.AgeInMonth,
		Description: p.Description,
		ImageUrls:   p.ImageUrls,
		UserId:      p.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &ResultCreate{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
	}, nil

}
