package catusecase

import "github.com/BennoAlif/ps-cats-social/src/entities"

func (i *sCatUsecase) Delete(catId *int) error {
	filters := entities.CatSearchFilter{
		ID: *catId,
	}

	cat, _ := i.catRepository.IsExists(&filters)

	if !cat {
		return ErrCatNotFound
	}

	err := i.catRepository.Delete(catId)

	if err != nil {
		return err
	}

	return nil
}
