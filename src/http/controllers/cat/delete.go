package catv1controller

import (
	"net/http"
	"strconv"

	catrepository "github.com/BennoAlif/ps-cats-social/src/repositories/cat"
	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
	catusecase "github.com/BennoAlif/ps-cats-social/src/usecase/cat"
	"github.com/labstack/echo/v4"
)

func (i *V1Cat) Delete(c echo.Context) (err error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'id'",
		})
	}

	uu := catusecase.New(
		catrepository.New(i.DB),
		userrepository.New(i.DB),
	)

	err = uu.Delete(&id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Cat deleted successfully",
		Data:    nil,
	})
}
