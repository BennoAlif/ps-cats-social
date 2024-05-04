package matchv1controller

import (
	"net/http"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	catrepository "github.com/BennoAlif/ps-cats-social/src/repositories/cat"
	matchrepository "github.com/BennoAlif/ps-cats-social/src/repositories/match"
	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
	matchusecase "github.com/BennoAlif/ps-cats-social/src/usecase/match"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func (i *V1Match) Delete(c echo.Context) (err error) {
	uid := new(meValidator)
	mapstructure.Decode(c.Get("user"), &uid)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'id'",
		})
	}

	uu := matchusecase.New(
		matchrepository.New(i.DB),
		catrepository.New(i.DB),
		userrepository.New(i.DB),
	)

	err = uu.Delete(&id, &uid.ID)

	if err != nil {
		if err.Error() == matchusecase.ErrForbidden.Error() {
			return c.JSON(http.StatusForbidden, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		} else {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "Cat deleted successfully",
		Data:    nil,
	})
}
