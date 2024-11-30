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
	claimsData := &model.Claims{}

	tkn, err := jwt.Parse(tknStr, func(token *jwt.Token) (interface{}, error) {
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

	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		if _, ok := claims["user_id"]; ok {
			claimsData.UserID = claims["user_id"].(string)
		}

		if _, ok := claims["email"]; ok {
			claimsData.Email = claims["email"].(string)
		}

		if _, ok := claims["exp"]; ok {
			claimsData.ExpiresAt = int64(claims["exp"].(float64))
		}
	} else {
		return nil, errors.New("invalid token")
	}

	return claimsData, nil
}
