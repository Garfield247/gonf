package utils

import (
	"time"

	"github.com/Garfield247/gonf.git/config"
	"github.com/dgrijalva/jwt-go/v4"
)

var jwtSecret = []byte(config.JwtCfg.SecterKey)

type Claims struct {
	UId      uint64 `json:"uid"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(uid uint64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(config.JwtCfg.LifeLength * time.Hour)

	claims := Claims{
		uid,
		username,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(expireTime),
			Issuer:    "bf1nf",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//  解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) { return jwtSecret, nil },
	)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
