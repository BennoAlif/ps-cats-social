package catrepository

import (
	"log"
	"time"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sCatRepository) Update(catId *int, p *entities.ParamsUpdateCat) (*entities.CreateCat, error) {
	var id int64
	var createdAt time.Time

	err := i.DB.QueryRow("UPDATE cats SET name = $2, race = $3, sex = $4, age_in_month = $5, description = $6, img_urls = $7 WHERE id = $1 RETURNING id, created_at",
		catId,
		p.Name,
		p.Race,
		p.Sex,
		p.AgeInMonth,
		p.Description,
		p.ImageUrls,
	).Scan(&id, &createdAt)

	if err != nil {
		log.Printf("Error updating cat: %s", err)
		return nil, err
	}

	cat := &entities.CreateCat{
		ID:        id,
		CreatedAt: createdAt,
	}

	return cat, nil
}
