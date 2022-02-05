package auth

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var signingKey = []byte(os.Getenv("JWTKEY"))

func CreateToken(userID uint) (string, error) {

	type customClaim struct {
		UserID uint `json:"id"`
		jwt.StandardClaims
	}
	// Create the Claims
	claims := customClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func DecodeToken(TokenString string) (uint, error) {
	if TokenString == "" {
		return 0, errors.New("cannot decode empty token")
	}

	token, err := jwt.Parse(TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := fmt.Sprintf("%v", claims["id"])
		if id == "" {
			return 0, fmt.Errorf("could not convert id to string. id value: %v", claims["id"])
		}
		parsed, err := strconv.ParseUint(id, 10, 64)
		return uint(parsed), err
	} else {
		return 0, err
	}
}
