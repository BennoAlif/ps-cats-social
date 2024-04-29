package user

import (
	"log"

	"github.com/BennoAlif/ps-cats-social/src/helper"
)

type Handler struct{}

func Create(user User) User {
	db := helper.CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO users (full_name, email, password) VALUES ($1, $2, $3) RETURNING full_name, email`

	var u User

	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&u.Name, &u.Email)

	if err != nil {
		log.Printf("Tidak Bisa mengeksekusi query. %v", err)
	}

	return u
}

func FindAll() ([]User, error) {
	db := helper.CreateConnection()
	defer db.Close()

	var users []User

	sqlStatement := `SELECT id, full_name, email, password FROM users`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		users = append(users, user)

	}
	return users, err
}
