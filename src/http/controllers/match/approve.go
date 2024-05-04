package matchv1controller

import (
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	catrepository "github.com/BennoAlif/ps-cats-social/src/repositories/cat"
	matchrepository "github.com/BennoAlif/ps-cats-social/src/repositories/match"
	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
	matchusecase "github.com/BennoAlif/ps-cats-social/src/usecase/match"
	"github.com/labstack/echo/v4"
)

func (i *V1Match) Approve(c echo.Context) (err error) {
	u := new(entities.SearchMatch)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

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

	err = uu.Approve(&u.MatchCatId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "Cat matching approved",
		Data:    nil,
	})
}
