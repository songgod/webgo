package common

import (
	"back/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtkey = []byte("a_secret_crect")

type Claims struct {
	jwt.StandardClaims
	UserId uint
}

func ReleaseToken(user model.User) (string, error) {
	expiredTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gptoil.com",
			Subject:   "user token test",
		},
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := tok.SignedString(jwtkey)
	if err != nil {
		return s, err
	}

	return s, nil
}

func ParseToken(tokenstring string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	// 从tokenstring里解析claims，返回
	token, err := jwt.ParseWithClaims(tokenstring, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})

	return token, claims, err
}
