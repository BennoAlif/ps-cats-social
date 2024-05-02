package catv1controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	catrepository "github.com/BennoAlif/ps-cats-social/src/repositories/cat"
	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
	catusecase "github.com/BennoAlif/ps-cats-social/src/usecase/cat"
	"github.com/labstack/echo/v4"
)

func (i *V1Cat) FindMany(c echo.Context) (err error) {
	filters := &entities.CatSearchFilter{}

	if idStr := c.QueryParam("id"); idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'id'",
			})
		}
		filters.ID = id
	}
	if limitStr := c.QueryParam("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'limit'",
			})
		}
		filters.Limit = limit
	}
	if offsetStr := c.QueryParam("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'offset'",
			})
		}
		filters.Offset = offset
	}
	if race := c.QueryParam("race"); race != "" {
		if !isValidRace(race) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'race'",
			})
		}
		filters.Race = race
	}
	if sex := c.QueryParam("sex"); sex != "" {
		if !isValidSex(sex) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'sex'",
			})
		}
		filters.Sex = sex
	}
	if hasMatchedStr := c.QueryParam("hasMatched"); hasMatchedStr != "" {
		hasMatched, err := strconv.ParseBool(hasMatchedStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'hasMatched'",
			})
		}
		filters.HasMatched = hasMatched
	}
	if ageInMonth := c.QueryParam("ageInMonth"); ageInMonth != "" {
		if !isValidAgeInMonth(ageInMonth) {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'ageInMonth'",
			})
		}
		filters.AgeInMonth = ageInMonth
	}
	if ownedStr := c.QueryParam("owned"); ownedStr != "" {
		owned, err := strconv.ParseBool(ownedStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'owned'",
			})
		}
		filters.Owned = owned
	}
	if search := c.QueryParam("search"); search != "" {
		filters.Search = search
	}

	uu := catusecase.New(
		catrepository.New(i.DB),
		userrepository.New(i.DB),
	)

	data, err := uu.FindMany(filters)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	// Return success response with data
	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Cats found successfully",
		Data:    data,
	})
}

// Helper functions for validation

func isValidRace(race string) bool {
	validRaces := []string{"Persian", "Maine Coon", "Siamese", "Ragdoll", "Bengal", "Sphynx", "British Shorthair", "Abyssinian", "Scottish Fold", "Birman"}
	for _, validRace := range validRaces {
		if race == validRace {
			return true
		}
	}
	return false
}

func isValidSex(sex string) bool {
	return sex == "male" || sex == "female"
}

func isValidAgeInMonth(ageInMonth string) bool {
	_, err := strconv.Atoi(ageInMonth)
	if err != nil {
		if strings.HasPrefix(ageInMonth, ">") || strings.HasPrefix(ageInMonth, "<") || strings.HasPrefix(ageInMonth, "=") {
			return true
		}
		return false
	}
	return true
}
