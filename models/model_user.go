package models

type ResponseUser struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type User struct {
	UserID        int `json:"user_id" form:"user_id"`
	UserFirstname string `json:"user_firstname" form:"user_firstname"`
}
