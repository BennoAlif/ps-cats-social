package matchv1controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1Match struct {
	DB *sql.DB
}

type iV1Match interface {
	Create(c echo.Context) error
	FindMany(c echo.Context) error
	Delete(c echo.Context) error
}

func New(v1Match *V1Match) iV1Match {
	return v1Match
}
