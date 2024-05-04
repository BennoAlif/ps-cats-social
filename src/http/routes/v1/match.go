package v1routes

import (
	matchV1Controller "github.com/BennoAlif/ps-cats-social/src/http/controllers/match"
	"github.com/BennoAlif/ps-cats-social/src/http/middlewares"
)

func (i *V1Routes) MountMatch() {
	g := i.Echo.Group("/cat/match")

	matchController := matchV1Controller.New(&matchV1Controller.V1Match{
		DB: i.DB,
	})

	g.POST("", matchController.Create, middlewares.Authentication())
	// g.GET("", matchController.FindMany, middlewares.Authentication())
	// g.PUT("/:id", matchController.Update, middlewares.Authentication())
	// g.DELETE("/:id", matchController.Delete, middlewares.Authentication())
}
