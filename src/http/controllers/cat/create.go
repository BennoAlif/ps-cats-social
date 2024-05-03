package catv1controller

import (
	"net/http"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	catrepository "github.com/BennoAlif/ps-cats-social/src/repositories/cat"
	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
	catusecase "github.com/BennoAlif/ps-cats-social/src/usecase/cat"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type (
	meValidator struct {
		ID int `mapstructure:"user_id" validate:"required"`
	}
)

func (i *V1Cat) Create(c echo.Context) (err error) {
	u := new(createRequest)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	uid := new(meValidator)
	mapstructure.Decode(c.Get("user"), &uid)

	if !ValidateRace(u.Race) {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Invalid race",
		})
	}

	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	uu := catusecase.New(
		catrepository.New(i.DB),
		userrepository.New(i.DB),
	)

	data, err := uu.Create(&entities.ParamsCreateCat{
		Name:        u.Name,
		Race:        u.Race,
		Sex:         u.Sex,
		AgeInMonth:  u.AgeInMonth,
		Description: u.Description,
		ImageUrls:   u.ImageUrls,
		UserId:      uid.ID,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Cat created successfully",
		Data:    data,
	})
}

type (
	createRequest struct {
		Name        string   `json:"name" validate:"required,min=1,max=30"`
		Race        string   `json:"race" validate:"required"`
		Sex         string   `json:"sex" validate:"required,oneof=male female"`
		AgeInMonth  int      `json:"ageInMonth" validate:"required,min=1,max=120082"`
		Description string   `json:"description" validate:"required,min=1,max=200"`
		ImageUrls   []string `json:"imageUrls" validate:"required,min=1,dive,url"`
	}
)

func ValidateRace(race string) bool {
	validRaces := map[string]bool{
		"Persian":           true,
		"Maine Coon":        true,
		"Siamese":           true,
		"Ragdoll":           true,
		"Bengal":            true,
		"Sphynx":            true,
		"British Shorthair": true,
		"Abyssinian":        true,
		"Scottish Fold":     true,
		"Birman":            true,
	}

	_, ok := validRaces[race]
	return ok
}
