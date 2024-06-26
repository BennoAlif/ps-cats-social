package catrepository

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sCatRepository) FindMany(filters *entities.CatSearchFilter) ([]*entities.Cat, error) {
	query := "SELECT id, name, race, sex, age_in_month, description, img_urls, created_at, user_id FROM cats WHERE 1=1 "
	params := []interface{}{}

	n := (&entities.CatSearchFilter{})

	if !reflect.DeepEqual(filters, n) {
		conditions := []string{}

		if filters.ID != 0 {
			conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.ID)
		}
		if filters.Search != "" {
			conditions = append(conditions, "name = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.Search)
		}
		if filters.Race != "" {
			conditions = append(conditions, "race = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.Race)
		}
		if filters.Sex != "" {
			conditions = append(conditions, "sex = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.Sex)
		}
		// if filters.HasMatched {
		// TODO: FIX THIS LOGIC
		// conditions = append(conditions, "matching_cat_id IS NOT NULL AND is_approved = true"+strconv.Itoa(len(params)+1))
		// params = append(params, filters.HasMatched)
		// }
		if filters.AgeInMonth != "" {
			operator, value := parseAgeInMonth(filters.AgeInMonth)
			conditions = append(conditions, fmt.Sprintf("age_in_month %s $%d", operator, len(params)+1))
			params = append(params, value)
		}
		if filters.Owned {
			conditions = append(conditions, "user_id = $"+strconv.Itoa(len(params)+1))
			params = append(params, filters.UserId)
		}

		if len(conditions) > 0 {
			query += " AND "
		}
		query += strings.Join(conditions, " AND ")
	}

	if filters.Limit == 0 {
		filters.Limit = 5
	}

	query += " ORDER BY created_at DESC"

	query += " LIMIT $" + strconv.Itoa(len(params)+1)
	params = append(params, filters.Limit)

	if filters.Offset == 0 {
		filters.Offset = 0
	} else {
		query += " OFFSET $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Offset)
	}

	rows, err := i.DB.Query(query, params...)
	if err != nil {
		log.Printf("Error finding cat: %s", err)
		return nil, err
	}
	defer rows.Close()

	cats := make([]*entities.Cat, 0)
	for rows.Next() {
		c := new(entities.Cat)
		err := rows.Scan(&c.ID, &c.Name, &c.Race, &c.Sex, &c.AgeInMonth, &c.Description, &c.ImageUrls, &c.CreatedAt, &c.UserId)
		if err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cats, nil
}

func parseAgeInMonth(ageInMonth string) (string, int) {
	value, err := strconv.Atoi(ageInMonth)
	if err != nil {
		operator := ageInMonth[:1]
		value, _ := strconv.Atoi(ageInMonth[1:])
		return operator, value
	}
	return "=", value
}
