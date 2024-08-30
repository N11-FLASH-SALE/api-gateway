package auth

import (
	"api/config"

	"github.com/dgrijalva/jwt-go"
)

var (
	acceskey = config.Load().ACCES_KEY
)

func ValidateAccessToken(tokenStr string) (bool, error) {
	_, err := ExtractAccessClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractAccessClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(acceskey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserInfoFromAccessToken(accessTokenString string) (string, string, error) {
	claims, err := ExtractAccessClaim(accessTokenString)
	if err != nil {
		return "", "", err
	}

	userID := (*claims)["user_id"].(string)
	Role := (*claims)["role"].(string)

	return userID, Role, nil
}
