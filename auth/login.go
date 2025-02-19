package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"squarepos/database"
	"squarepos/dto"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter,req *http.Request){
	var LoginUser dto.LoginReq
	err:=json.NewDecoder(req.Body).Decode(&LoginUser)
	if err!=nil{
		fmt.Println("error in decoding Login credentials")
	}
	fmt.Println("added",LoginUser,LoginUser.Username)

	var RetrivedUser dto.User
	query := `SELECT id, username, password, access_token FROM restaurant_users WHERE username = $1`
	err = database.Db.QueryRow(query, LoginUser.Username).Scan(&RetrivedUser.UserID, &RetrivedUser.Username, &RetrivedUser.Password, &RetrivedUser.AccessToken)
	if err==sql.ErrNoRows{
		fmt.Println("no users found ")
		return
	}
	if err!=nil{
		fmt.Print("error in query ", err)
	}
	fmt.Println("ret",RetrivedUser)

	err=bcrypt.CompareHashAndPassword([]byte(RetrivedUser.Password),[]byte(LoginUser.Password))
	if err!=nil{
		fmt.Println("Invalid credentials")
		return
	}

	token,err:=GenarateToken(RetrivedUser.UserID,RetrivedUser.AccessToken)
	if err!=nil{
		fmt.Print("error in generating tokens")
	}
	json.NewEncoder(w).Encode(map[string]string{
		"token":token,
	})
}