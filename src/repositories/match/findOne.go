package matchrepository

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sMatchRepository) FindOne(filters *entities.SearchMatch) (*entities.FindOneMatch, error) {
	query := "SELECT id, user_cat_id, match_cat_id, status, message, created_at, updated_at FROM cat_matches WHERE "
	params := []interface{}{}
	conditions := []string{}

	if filters.MatchCatId != 0 {
		conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filters.MatchCatId)
	}

	query += strings.Join(conditions, " AND ")

	query += " LIMIT 1"

	row := i.DB.QueryRow(query, params...)

	var match entities.FindOneMatch
	err := row.Scan(
		&match.ID,
		&match.UserCatId,
		&match.MatchCatId,
		&match.Status,
		&match.Message,
		&match.CreatedAt,
		&match.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil for both user and error
		}
		return nil, err
	}
	return &match, nil
}
