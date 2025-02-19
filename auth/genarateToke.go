package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct{
	UserID      int    `json:"user_id"`
	AccessToken string `json:"access_token"`
	jwt.RegisteredClaims


}

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}

	return value
}

var JwtSecret = []byte(GetEnv("JWTSECRET"))

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


