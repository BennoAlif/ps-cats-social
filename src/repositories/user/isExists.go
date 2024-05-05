package userrepository

import (
	"log"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

func (i *sUserRepository) IsExists(filters *entities.ParamsCreateUser) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE "
	params := []interface{}{}
	conditions := []string{}

	if filters.ID != 0 {
		conditions = append(conditions, "id = $"+strconv.Itoa(len(params)+1))
		params = append(params, filters.ID)
	}
	if filters.Name != "" {
		conditions = append(conditions, "name = $"+strconv.Itoa(len(params)+1))
		params = append(params, filters.Name)
	}
	if filters.Email != "" {
		conditions = append(conditions, "email = $"+strconv.Itoa(len(params)+1))
		params = append(params, filters.Email)
	}

	query += strings.Join(conditions, " AND ")
	query += ")"

	var exists bool
	err := i.DB.QueryRow(query, params...).Scan(&exists)

	if err != nil {
		log.Printf("Error checking if user exists: %s", err)
		return false, err
	}

	return exists, nil
}
