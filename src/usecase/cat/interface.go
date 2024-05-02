package catusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
	cat "github.com/BennoAlif/ps-cats-social/src/repositories/cat"
	user "github.com/BennoAlif/ps-cats-social/src/repositories/user"
)

type sCatUsecase struct {
	catRepository  cat.CatRepository
	userRepository user.UserRepository
}

type CatUsecase interface {
	Create(*entities.ParamsCreateCat) (*ResultCreate, error)
	FindMany(*entities.CatSearchFilter) ([]*entities.Cat, error)
	Update(*int, *entities.ParamsUpdateCat) (*ResultCreate, error)
	Delete(*int) error
}

func New(catRepository cat.CatRepository, userRepository user.UserRepository) CatUsecase {
	return &sCatUsecase{
		catRepository:  catRepository,
		userRepository: userRepository,
	}
}
