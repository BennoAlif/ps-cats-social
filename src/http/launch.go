package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BennoAlif/ps-cats-social/src/helpers"
	v1routes "github.com/BennoAlif/ps-cats-social/src/http/routes/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func (i *Http) Launch() {
	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = helpers.ErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	basePath := "/"
	if os.Getenv("BASE_PATH") != "" {
		basePath = os.Getenv("BASE_PATH")
	}

	baseUrl := e.Group(basePath + "/v1")
	baseUrl.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf("API Base Code for %s", os.Getenv("ENVIRONMENT")))
	})

	baseUrl.Static("/static", "../../temp_assets")

	v1 := v1routes.New(&v1routes.V1Routes{
		Echo: e.Group(basePath + "/v1"),
		DB:   i.DB,
	})

	v1.MountPing()
	v1.MountUser()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))))
}
