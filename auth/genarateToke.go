package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

//later take from env
var JwtSecret = []byte("Abvslkfssjfdsjfdn1111121212")

type Claims struct{
	UserID      int    `json:"user_id"`
	AccessToken string `json:"access_token"`
	jwt.RegisteredClaims


}

func GenarateToken(UserId int,AcessToken string) (string,error){
	claims:=Claims{
		UserID: UserId,
		AccessToken: AcessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour*24)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(JwtSecret)
}


