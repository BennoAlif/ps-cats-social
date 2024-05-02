package v1controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type (
	meValidator struct {
		ID int `mapstructure:"user_id" validate:"required"`
	}
)

func (i *V1) PingAuth(c echo.Context) error {

	u := new(meValidator)
	mapstructure.Decode(c.Get("user"), &u)

	return c.JSON(http.StatusOK, Response{
		Status:  true,
		Message: "OK",
		Data:    "PONG AUTH",
	})
}
