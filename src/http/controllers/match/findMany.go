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

func (i *V1Match) FindMany(c echo.Context) (err error) {

	uid := new(meValidator)
	mapstructure.Decode(c.Get("user"), &uid)

	uu := matchusecase.New(
		matchrepository.New(i.DB),
		catrepository.New(i.DB),
		userrepository.New(i.DB),
	)

	data, err := uu.FindMany(&uid.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "Cats found successfully",
		Data:    data,
	})
}
