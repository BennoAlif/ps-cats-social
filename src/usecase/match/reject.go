package matchusecase

import "github.com/BennoAlif/ps-cats-social/src/entities"

func (i *sMatchUsecase) Reject(id *int) error {
	filters := entities.SearchMatch{
		MatchCatId: *id,
	}

	match, _ := i.matchRepository.IsExists(&filters)

	if !match {
		return ErrNotFound
	}

	err := i.matchRepository.Reject(id)

	if err != nil {
		return err
	}

	return nil
}
