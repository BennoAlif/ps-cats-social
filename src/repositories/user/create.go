package userrepository

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/entities"
)

type (
	ParamsCreateUser struct {
		Name     string
		Email    string
		Password string
	}
)

func (i *sUserRepository) Create(p *ParamsCreateUser) (*entities.User, error) {

	result, err := i.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", p.Name, p.Email, p.Password)
	if err != nil {
		log.Fatal(err)
	}

	// Get the ID of the inserted user
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	user := &entities.User{
		ID:    id,
		Name:  p.Name,
		Email: p.Email,
	}

	return user, nil

}
