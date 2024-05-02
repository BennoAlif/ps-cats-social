package v1routes

import (
	catV1Controller "github.com/BennoAlif/ps-cats-social/src/http/controllers/cat"
	"github.com/BennoAlif/ps-cats-social/src/http/middlewares"
)

func (i *V1Routes) MountCat() {
	g := i.Echo.Group("/cat")

	catController := catV1Controller.New(&catV1Controller.V1Cat{
		DB: i.DB,
	})

	g.POST("", catController.Create, middlewares.Authentication())
	g.GET("", catController.FindMany, middlewares.Authentication())
	g.PUT("/:id", catController.Update, middlewares.Authentication())
	g.DELETE("", catController.Delete, middlewares.Authentication())
}
