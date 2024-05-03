package userusecase

import (
	"os"

	"github.com/BennoAlif/ps-cats-social/src/entities"
	"github.com/BennoAlif/ps-cats-social/src/helpers"
)

type (
	ParamsLogin struct {
		Email    string
		Password string
	}
	GeneratedToken struct {
		Token     string `json:"token"`
		ExpiredAt int64  `json:"expired_at"`
	}
	ResultLogin struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		AccessToken string `json:"accessToken"`
	}
)

func (i *sUserUsecase) Login(p *ParamsLogin) (*ResultLogin, error) {

	filters := entities.ParamsCreateUser{
		Email: p.Email,
	}

	user, _ := i.userRepository.FindOne(&filters)

	if user == nil {
		return nil, ErrUserNotFound
	}

	paramsGenerateJWTLogin := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 480,
		UserId:          user.ID,
		SecretKey:       os.Getenv("JWT_SECRET"),
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, user.Password)
	if !isValidPassword {
		return nil, ErrInvalidPassword
	}

	accessToken, _, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTLogin)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	return &ResultLogin{
		Name:        user.Name,
		Email:       p.Email,
		AccessToken: accessToken,
	}, nil
}
