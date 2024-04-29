package user

import (
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct{}

func (c *Controller) FindAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	users, err := FindAll()
	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}
	json.NewEncoder(w).Encode(users)
}
