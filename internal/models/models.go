package models

type User struct {
	Id       int
	Username string
	Password string
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
