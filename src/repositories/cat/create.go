package catrepository

import (
	"log"
	"time"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sCatRepository) Create(p *entities.ParamsCreateCat) (*entities.CreateCat, error) {
	var id int64
	var createdAt time.Time
	err := i.DB.QueryRow("INSERT INTO cats (name, race, sex, age_in_month, description, img_urls, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at",
		p.Name,
		p.Race,
		p.Sex,
		p.AgeInMonth,
		p.Description,
		p.ImageUrls,
		p.UserId,
	).Scan(&id, &createdAt)

	if err != nil {
		log.Printf("Error inserting cat: %s", err)
		return nil, err
	}

	cat := &entities.CreateCat{
		ID:        id,
		CreatedAt: createdAt,
	}

	return cat, nil
}
