package catrepository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sCatRepository) IsExists(filters *entities.CatSearchFilter) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM cats WHERE 1=1 "
	params := []interface{}{}

	if filters.ID != 0 {
		query += "AND id = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.ID)
	}
	if filters.Search != "" {
		query += "AND name = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Search)
	}
	if filters.Race != "" {
		query += "AND race = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Race)
	}
	if filters.Sex != "" {
		query += "AND sex = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.Sex)
	}
	if filters.AgeInMonth != "" {
		operator, value := parseAgeInMonth(filters.AgeInMonth)
		query += fmt.Sprintf("AND age_in_month %s $%d ", operator, len(params)+1)
		params = append(params, value)
	}
	if filters.Owned {
		query += "AND user_id = $" + strconv.Itoa(len(params)+1)
		params = append(params, filters.UserId)
	}

	query += ")"

	var exists bool
	err := i.DB.QueryRow(query, params...).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if cat exists: %s", err)
		return false, err
	}

	return exists, nil
}
