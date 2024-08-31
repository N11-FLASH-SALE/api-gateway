package auth

import (
	"api/config"

	"github.com/dgrijalva/jwt-go"
)

var (
	refresh = config.Load().REFRESH_KEY
)

func ValidateRefreshToken(tokenStr string) (bool, error) {
	_, err := ExtractRefreshClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractRefreshClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(refresh), nil
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

func GetUserInfoFromRefreshToken(refreshTokenString string) (string, string, error) {
	claims, err := ExtractRefreshClaim(refreshTokenString)
	if err != nil {
		return "", "", err
	}

	userID := (*claims)["user_id"].(string)
	Role := (*claims)["role"].(string)

	return userID, Role, nil
}
