package matchrepository

import (
	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sMatchRepository) Create(p *entities.CreateMatch) error {
	_, err := i.DB.Exec("INSERT INTO cat_matches (user_cat_id, match_cat_id, message) VALUES ($1, $2, $3)",
		p.UserCatId,
		p.MatchCatId,
		p.Message,
	)

	if err != nil {
		return err
	}

	return nil
}
