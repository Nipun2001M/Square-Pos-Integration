package dto

type User struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Password    string `json:"password"` 
	AccessToken string `json:"access_token"`
}

type LoginReq struct{
	Username    string `json:"username"`
	Password    string `json:"password"`

}
