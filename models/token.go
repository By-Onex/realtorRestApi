package models

import "github.com/dgrijalva/jwt-go"

type UserT string

//Token модель
type Token struct {
	UserID int64
	jwt.StandardClaims
}
