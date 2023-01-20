package util

import (
	"github.com/itachilee/ginblog/global"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(global.AppSetting.JwtSecret)

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Sub      string
	jwt.StandardClaims
}

func GenerateToken(id int, username, sub string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(global.JwtSetting.ExpiresSpan) * time.Second)

	Claims := Claims{
		id,
		username,
		sub,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JwtSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)

	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil

	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}

	}
	return nil, err
}
