package models

type User struct {
	ID       int64
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Repassword string `json:"repeat_password"`
}

type Session struct {
	ID          int
	UserID      int
	Token       string
	ExpiredDate string
}
