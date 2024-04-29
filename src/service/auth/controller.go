package auth

import (
	"encoding/json"
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/helper"
	"github.com/BennoAlif/ps-cats-social/src/service/user"
	"github.com/go-playground/validator/v10"
)

type Controller struct{}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var dto user.User
	err := helper.ReadRequest(&dto, r)

	if err != nil {
		w.Write([]byte("error"))
	}

	validate := validator.New()
	err = validate.Struct(dto)

	if err != nil {
		w.Write([]byte(err.Error()))

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	createdUser := user.Create(dto)

	json.NewEncoder(w).Encode(createdUser)

}
