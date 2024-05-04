package matchusecase

import "github.com/BennoAlif/ps-cats-social/src/entities"

func (i *sMatchUsecase) Approve(id *int) error {
	filters := entities.SearchMatch{
		MatchCatId: *id,
	}

	match, _ := i.matchRepository.FindOne(&filters)

	if match == nil {
		return ErrNotFound
	}

	err := i.matchRepository.Approve(id)

	if err != nil {
		return err
	}

	return nil
}
