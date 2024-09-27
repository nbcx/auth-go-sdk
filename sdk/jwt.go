package sdk

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	User
	AccessToken string `json:"accessToken"`
	jwt.RegisteredClaims
	TokenType        string `json:"tokenType"`
	RefreshTokenType string `json:"TokenType"`
}

// IsRefreshToken returns true if the token is a refresh token
func (c Claims) IsRefreshToken() bool {
	return c.RefreshTokenType == "refresh-token"
}

func (c *Client) ParseJwtToken(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(c.Certificate))
		if err != nil {
			return nil, err
		}

		return publicKey, nil
	})

	if t != nil {
		if claims, ok := t.Claims.(*Claims); ok && t.Valid {
			return claims, nil
		}
	}

	return nil, err
}
