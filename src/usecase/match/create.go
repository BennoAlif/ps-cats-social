package matchusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
	catusecase "github.com/BennoAlif/ps-cats-social/src/usecase/cat"
)

func (i *sMatchUsecase) Create(p *entities.CreateMatch, uId *int) error {

	matchCatId, _ := i.catRepository.FindMany(&entities.CatSearchFilter{
		ID:    p.MatchCatId,
		Limit: 1,
	})

	if len(matchCatId) == 0 {
		return catusecase.ErrCatNotFound
	}

	userCatId, _ := i.catRepository.FindMany(&entities.CatSearchFilter{
		ID:    p.UserCatId,
		Limit: 1,
	})

	if len(userCatId) == 0 {
		return catusecase.ErrCatNotFound
	}

	// if userCatId[0].UserId.Valid && int(userCatId[0].UserId.Int64) != *uId {
	// 	return ErrCatNotBelongToUser
	// }

	if matchCatId[0].Sex == userCatId[0].Sex {
		return ErrSameCatGender
	}

	// if matchCatId[0].IsApproved.Valid && matchCatId[0].IsApproved.Bool {
	// 	return ErrAlreadyMatch
	// }

	// if userCatId[0].IsApproved.Valid && userCatId[0].IsApproved.Bool {
	// 	return ErrAlreadyMatch
	// }

	// if userCatId[0].UserId == matchCatId[0].UserId {
	// 	return ErrSameOwner
	// }

	err := i.matchRepository.Create(&entities.CreateMatch{
		MatchCatId: p.MatchCatId,
		UserCatId:  p.UserCatId,
		Message:    p.Message,
	})

	if err != nil {
		return err
	}

	return nil
}
