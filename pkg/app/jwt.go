package app

import (
	"github.com/golang-jwt/jwt"
	"github.com/itachilee/ginblog/global"
	"time"
)

var jwtSecret = []byte(global.JWTSetting.Secret)

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id int, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(global.JWTSetting.Expire) * time.Second)

	Claims := Claims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
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
