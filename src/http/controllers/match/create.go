package matchv1controller

import (
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	catrepository "github.com/BennoAlif/ps-cats-social/src/repositories/cat"
	matchrepository "github.com/BennoAlif/ps-cats-social/src/repositories/match"
	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
	matchusecase "github.com/BennoAlif/ps-cats-social/src/usecase/match"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type (
	meValidator struct {
		ID int `mapstructure:"user_id" validate:"required"`
	}
)

func (i *V1Match) Create(c echo.Context) (err error) {
	u := new(entities.CreateMatch)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	uid := new(meValidator)
	mapstructure.Decode(c.Get("user"), &uid)

	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	uu := matchusecase.New(
		matchrepository.New(i.DB),
		catrepository.New(i.DB),
		userrepository.New(i.DB),
	)

	err = uu.Create(&entities.CreateMatch{
		MatchCatId: u.MatchCatId,
		UserCatId:  u.UserCatId,
		Message:    u.Message,
	}, &uid.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entities.SuccessResponse{
		Message: "Cat matched successfully",
		Data:    nil,
	})
}
