package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"squarepos/database"
	"squarepos/dto"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter,req *http.Request){
	var LoginUser dto.LoginReq
	err:=json.NewDecoder(req.Body).Decode(&LoginUser)
	if err!=nil{
		http.Error(w, "error in decoding credentials ", http.StatusBadRequest)
		return
	}
	var RetrivedUser dto.User
	query := `SELECT id, username, password, access_token FROM restaurant_users WHERE username = $1`
	err = database.Db.QueryRow(query, LoginUser.Username).Scan(&RetrivedUser.UserID, &RetrivedUser.Username, &RetrivedUser.Password, &RetrivedUser.AccessToken)
	if err==sql.ErrNoRows{
		http.Error(w, "No Users Found", http.StatusBadRequest)
		return
	}
	if err!=nil{
		http.Error(w, "error in query ", http.StatusBadRequest)
		return
	}
	err=bcrypt.CompareHashAndPassword([]byte(RetrivedUser.Password),[]byte(LoginUser.Password))
	if err!=nil{
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	token,err:=GenarateToken(RetrivedUser.UserID,RetrivedUser.AccessToken)
	if err!=nil{
		http.Error(w, "Error in Genereating new Token", http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(map[string]string{
		"token":token,
	})
}