package libs

import (
	"fmt"

	"github.com/AlejandroAldana99/mvp_api/constants"
	"github.com/golang-jwt/jwt/v4"
)

func DecodeJWT(token string) (string, error) {
	hmacSecret := []byte("secret")

	tokenD, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if claims, ok := tokenD.Claims.(jwt.MapClaims); ok {
		return claims["role"].(string), nil
	} else {
		return constants.GenericName, err
	}
}
