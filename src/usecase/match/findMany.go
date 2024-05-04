package matchusecase

import "github.com/BennoAlif/ps-cats-social/src/entities"

func (i *sMatchUsecase) FindMany(uId *int) ([]*entities.Match, error) {
	matches, err := i.matchRepository.FindMany(uId)

	if err != nil {
		return nil, err
	}

	return matches, nil
}
