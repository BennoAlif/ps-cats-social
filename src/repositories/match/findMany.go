package matchrepository

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sMatchRepository) FindMany(uId *int) ([]*entities.Match, error) {
	query := `SELECT 
m.id AS id,
u.name AS issued_name,
u.email AS issued_email,
u.created_at AS issued_created_at,
c.id AS issued_cat_id,
c.name AS issued_cat_name,
c.race AS issued_cat_race,
c.sex AS issued_cat_sex,
c.description AS issued_cat_description,
c.age_in_month AS issued_cat_age_in_month,
c.img_urls AS issued_cat_img_urls,
CASE 
			WHEN m.status = 'approved' THEN true
			ELSE false
	END AS issued_has_matched,
c.created_at AS issued_cat_created_at,
mc.id AS match_cat_id,
mc.name AS match_cat_name,
mc.race AS match_cat_race,
mc.sex AS match_cat_sex,
mc.description AS match_cat_description,
mc.age_in_month AS match_cat_age_in_month,
mc.img_urls AS match_cat_img_urls,
CASE 
			WHEN m.status = 'approved' THEN true
			ELSE false
	END AS match_has_matched,
mc.created_at AS match_cat_created_at,
m.message AS message,
m.created_at AS created_at
FROM 
cat_matches m
JOIN
cats c
ON
c.id = m.user_cat_id
JOIN
users u
ON
u.id = c.user_id
JOIN
cats mc
ON
mc.id = m.match_cat_id
JOIN
users mu
ON
mu.id = mc.user_id
WHERE u.id = $1 OR mu.id = $1
ORDER BY id DESC
	`

	rows, err := i.DB.Query(query, uId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	matches := make([]*entities.Match, 0)
	for rows.Next() {
		m := new(entities.Match)
		err := rows.Scan(
			&m.ID,
			&m.IssuedBy.Name,
			&m.IssuedBy.Email,
			&m.IssuedBy.CreatedAt,
			&m.UserCatDetail.ID,
			&m.UserCatDetail.Name,
			&m.UserCatDetail.Race,
			&m.UserCatDetail.Sex,
			&m.UserCatDetail.Description,
			&m.UserCatDetail.AgeInMonth,
			&m.UserCatDetail.ImageUrls,
			&m.UserCatDetail.HasMatched,
			&m.UserCatDetail.CreatedAt,
			&m.MatchCatDetail.ID,
			&m.MatchCatDetail.Name,
			&m.MatchCatDetail.Race,
			&m.MatchCatDetail.Sex,
			&m.MatchCatDetail.Description,
			&m.MatchCatDetail.AgeInMonth,
			&m.MatchCatDetail.ImageUrls,
			&m.MatchCatDetail.HasMatched,
			&m.MatchCatDetail.CreatedAt,
			&m.Message,
			&m.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		matches = append(matches, m)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error finding match: %s", err)
		return nil, err
	}

	return matches, nil
}
