package userusecase

import (
	"os"
	"strconv"

	"github.com/BennoAlif/ps-cats-social/src/helpers"

	userrepository "github.com/BennoAlif/ps-cats-social/src/repositories/user"
)

type (
	ParamsCreateUser struct {
		Email    string
		Name     string
		Password string
	}
)

func (i *sUserUsecase) CreateUser(p *ParamsCreateUser) (*ResultLogin, error) {
	expiredInMinutesStr := os.Getenv("ACCESS_TOKEN_EXPIRED_IN_MINUTES")
	expiredInMinutes, _ := strconv.Atoi(expiredInMinutesStr)

	checkEmail, _ := i.userRepository.FindByEmail(&p.Email)

	if checkEmail != nil {
		return nil, ErrEmailAlreadyUsed
	}

	hashedPassword, _ := helpers.HashPassword(p.Password)
	data, err := i.userRepository.Create(&userrepository.ParamsCreateUser{
		Email:    p.Email,
		Name:     p.Name,
		Password: hashedPassword,
	})

	paramsGenerateJWTRegister := helpers.ParamsGenerateJWT{
		ExpiredInMinute: expiredInMinutes,
		UserId:          data.ID,
		SecretKey:       os.Getenv("JWT_SECRET"),
	}

	accessToken, _, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTRegister)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	if err != nil {
		return nil, err
	}

	return &ResultLogin{
		Name:        p.Name,
		Email:       p.Email,
		AccessToken: accessToken,
	}, nil
}
