package matchrepository

import (
	"log"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sMatchRepository) IsExists(filters *entities.SearchMatch) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM cat_matches WHERE "
	params := []interface{}{}
	conditions := []string{}

	if filters.MatchCatId != 0 {
		conditions = append(conditions, "match_cat_id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filters.MatchCatId)
	}

	query += strings.Join(conditions, " AND ")
	query += ")"

	var exists bool
	err := i.DB.QueryRow(query, params...).Scan(&exists)

	if err != nil {
		log.Printf("Error checking if match exists: %s", err)
		return false, err
	}

	return exists, nil
}
