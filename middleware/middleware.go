package middleware

import (
	"fmt"
	"net/http"
	"squarepos/auth"
	"strings"
    "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            tokenRecieve:=r.Header.Get("Authorization")
            if tokenRecieve==""{
                http.Error(w, "error in token recieve", http.StatusBadRequest)
                return
            }
            claim:=&auth.Claims{}
            tokenTrimed := strings.TrimPrefix(tokenRecieve, "Bearer ")
            token,err:=jwt.ParseWithClaims(tokenTrimed,claim,func(t *jwt.Token) (interface{}, error) {
                return auth.JwtSecret,nil
            })

            if err != nil || !token.Valid {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                fmt.Println("error is",err)
                return
            }
            next.ServeHTTP(w,r)

        },
    )

}