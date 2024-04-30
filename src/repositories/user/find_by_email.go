package userrepository

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (i *sUserRepository) FindByEmail(email *string) (*entities.User, error) {
	rows, err := i.DB.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user entities.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return &user, nil
}
