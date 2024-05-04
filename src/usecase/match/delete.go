package matchusecase

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sMatchUsecase) Delete(id *int, uid *int) error {
	filters := entities.SearchMatch{
		MatchCatId: *id,
	}

	match, _ := i.matchRepository.FindOne(&filters)

	if match == nil {
		return ErrNotFound
	}

	// userCatId, err := strconv.Atoi(match.UserCatId)

	// if err != nil {
	// 	return err
	// }

	// fmt.Println(userCatId)
	// fmt.Println(*uid)

	// if userCatId != *uid {
	// 	return ErrForbidden
	// }

	err := i.matchRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
