package catv1controller

import (
	"net/http"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	catrepository "github.com/BennoAlif/ps-cats-social/src/repositories/cat"
	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
	catusecase "github.com/BennoAlif/ps-cats-social/src/usecase/cat"
	"github.com/labstack/echo/v4"
)

func (i *V1Cat) Update(c echo.Context) (err error) {
	u := new(createRequest)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid value for 'id'",
		})
	}

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err = c.Validate(u); err != nil {
		return err
	}

	uu := catusecase.New(
		catrepository.New(i.DB),
		userrepository.New(i.DB),
	)

	data, err := uu.Update(&id, &entities.ParamsUpdateCat{
		Name:        u.Name,
		Race:        u.Race,
		Sex:         u.Sex,
		AgeInMonth:  u.AgeInMonth,
		Description: u.Description,
		ImageUrls:   u.ImageUrls,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Cat updated successfully",
		Data:    data,
	})
}
