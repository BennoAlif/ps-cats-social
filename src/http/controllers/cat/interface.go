package catv1controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1Cat struct {
	DB *sql.DB
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type iV1Cat interface {
	Create(c echo.Context) error
	FindMany(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func New(v1Cat *V1Cat) iV1Cat {
	return v1Cat
}
