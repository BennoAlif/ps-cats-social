package matchusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
	cat "github.com/BennoAlif/ps-cats-social/src/repositories/cat"
	match "github.com/BennoAlif/ps-cats-social/src/repositories/match"
	user "github.com/BennoAlif/ps-cats-social/src/repositories/user"
)

type sMatchUsecase struct {
	matchRepository match.MatchRepository
	catRepository   cat.CatRepository
	userRepository  user.UserRepository
}

type MatchUsecase interface {
	Create(*entities.CreateMatch, *int) error
}

func New(matchRepository match.MatchRepository, catRepository cat.CatRepository, userRepository user.UserRepository) MatchUsecase {
	return &sMatchUsecase{
		matchRepository: matchRepository,
		catRepository:   catRepository,
		userRepository:  userRepository,
	}
}
