package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func Protect(token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("==signature=="), nil
	})

	return err
}
