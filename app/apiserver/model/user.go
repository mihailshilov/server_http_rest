package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// User
type User1 struct {
	ID       int    `json:"-"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Validation data booking
func (d *User1) ValidateUser1() error {
	return validation.ValidateStruct(
		d,
		validation.Field(&d.Email, validation.Required),
		validation.Field(&d.Password, validation.Required),
	)
}

// for jwt verify
type User2 struct {
	UserID uint64
}

// for token and exp
type Token_exp struct {
	Token string    `extensions:"x-order=a"`
	Exp   time.Time `extensions:"x-order=b" example:"2023-06-11T10:18:29+03:00"`
}

type AccessDetails struct {
	UserId uint64
	Exp    uint64
}

// response struct
type Response struct {
	Status   string `json:"status"`
	Response string `json:"response"`
}

// response struct booking
type ResponseBooking struct {
	StatusMs       string `json:"status_ms"`
	ResponseMs     string `json:"response_ms"`
	StatusLK       string `json:"status_lk"`
	ResponseGazCrm string `json:"response_gcrm"`
}

type HTTPerrIncorrectEmailOrPassword struct {
	Error string `json:"error" example:"incorrect auth"`
}

type HTTPerrReg struct {
	Error string `json:"error" example:"service registration error"`
}

type HTTPerrJwt struct {
	Error string `json:"error" example:"token error"`
}

type HTTPerrFindUser struct {
	Error string `json:"error" example:"user not found"`
}

type HTTPerrMssql struct {
	Error string `json:"error" example:"mssql error"`
}
