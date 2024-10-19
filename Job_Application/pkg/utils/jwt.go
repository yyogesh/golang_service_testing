package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Clamis struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateToken(username string, userID int, isAdmin bool) (string, error) {
	exiprationTime := time.Now().Add(10 * time.Hour)

	claims := &Clamis{
		Username: username,
		UserID:   userID,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exiprationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(token string) (*Clamis, error) {
	claims := &Clamis{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		return nil, err
	}

	return claims, nil
}
