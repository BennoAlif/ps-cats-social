package userusecase

import (
	"os"
	"strconv"

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
		Email       string
		Name        string
		AccessToken GeneratedToken
	}
)

func (i *sUserUsecase) Login(p *ParamsLogin) (*ResultLogin, error) {
	expiredInMinutesStr := os.Getenv("ACCESS_TOKEN_EXPIRED_IN_MINUTES")
	expiredInMinutes, _ := strconv.Atoi(expiredInMinutesStr)

	emailMx := helpers.ValidateMx(p.Email)

	if emailMx != nil {
		return nil, emailMx
	}
	user, _ := i.userRepository.FindByEmail(&p.Email)

	if user == nil {
		return nil, ErrInvalidUser
	}

	paramsGenerateJWTLogin := helpers.ParamsGenerateJWT{
		ExpiredInMinute: expiredInMinutes,
		UserId:          user.ID,
		SecretKey:       os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, user.Password)
	if !isValidPassword {
		return nil, ErrInvalidUser
	}

	accessToken, _, errAccessToken := helpers.GenerateJWT(&paramsGenerateJWTLogin)

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	return &ResultLogin{
		Name:  user.Name,
		Email: p.Email,
		AccessToken: GeneratedToken{
			Token: accessToken,
		},
	}, nil
}
