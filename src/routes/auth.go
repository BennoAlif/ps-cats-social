package routes

import (
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/service/auth"
)

func LoadAuthRoutes(router *http.ServeMux) {
	authController := &auth.Controller{}

	router.HandleFunc("POST /user/register", authController.Register)
	// router.HandleFunc("POST /user/login", authController.)
}
