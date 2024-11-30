package helper

import (
	"errors"
	"os"
	"simple-dashboard-server/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = os.Getenv("JWT_SECRET")

func GenerateToken(userID, email string, keepSignIn bool) (string, error) {
	expirationTime := time.Now()

	if keepSignIn {
		expirationTime = expirationTime.Add(time.Hour * 24 * 7)
	} else {
		expirationTime = expirationTime.Add(time.Hour * 24)
	}

	claims := &model.Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tknStr string) (*model.Claims, error) {
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
