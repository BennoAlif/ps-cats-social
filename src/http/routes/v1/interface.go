package v1routes

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type V1Routes struct {
	Echo *echo.Group
	DB   *sql.DB
}

type iV1Routes interface {
	MountPing()
}

func New(v1Routes *V1Routes) iV1Routes {
	return v1Routes
}
