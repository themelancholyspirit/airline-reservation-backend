package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey" autoIncrement` // Changed to pointer type
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-" gorm:"column:password"`
}

type UserSignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
