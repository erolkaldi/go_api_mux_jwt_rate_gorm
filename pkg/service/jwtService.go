package service

import (
	"os"
	"strconv"
	"time"

	"github.com/erolkaldi/agency/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (models.TokenDto, error) {
	expiration := time.Now().Add(time.Hour * 7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      user.Email,
		"name":       user.Name,
		"id":         strconv.Itoa(user.ID),
		"expiration": expiration.Unix(),
	})
	secret := os.Getenv("JWT_SECRET")
	access_token, err := token.SignedString([]byte(secret))
	if err != nil {
		println("******" + err.Error() + "*******")
		return models.TokenDto{}, err
	}
	return models.TokenDto{Access_Token: access_token, Expiration: expiration}, nil
}
func ValidateToken(access_token string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
