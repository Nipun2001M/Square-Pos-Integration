package auth

import (
	"encoding/json"
	"net/http"
	"squarepos/database"
	"squarepos/dto"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter,req *http.Request){
	var user dto.User
	err:=json.NewDecoder(req.Body).Decode(&user)
	if err!=nil{
		http.Error(w, "error in decoding user in register", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err!=nil{
		http.Error(w, "error in bycrypting passwords", http.StatusBadRequest)
		return
	}
	query:=`INSERT INTO restaurant_users (username,password,access_token) VALUES ($1, $2, $3)`
	_, err = database.Db.Exec(query, user.Username, string(hashedPassword), user.AccessToken)
	if err!=nil{
		http.Error(w, "error in adding to database", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})




}


