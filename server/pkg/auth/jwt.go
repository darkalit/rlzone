package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/darkalit/rlzone/server/config"
)

type TokenType string

const (
	RefreshTokenType TokenType = "refresh_token_type"
	AccessTokenType  TokenType = "access_token_type"
)

type JWTPayload struct {
	UserID uint
	Email  string
	Role   string
}

func GenJWT(payload *JWTPayload, config *config.Config, tokenType TokenType) (string, error) {
	var expTime int
	var secret string

	switch tokenType {
	case RefreshTokenType:
		expTime = config.RefreshTokenExpiryHour
		secret = config.RefreshTokenSecret
		break
	case AccessTokenType:
		expTime = config.AccessTokenExpiryHour
		secret = config.AccessTokenSecret
		break
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   payload.UserID,
		"iss":   config.TokenIssuer,
		"aud":   payload.Role,
		"email": payload.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(expTime)).Unix(),
		"iat":   time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(
	tokenString string,
	config *config.Config,
	tokenType TokenType,
) (*JWTPayload, error) {
	var secret string

	switch tokenType {
	case RefreshTokenType:
		secret = config.RefreshTokenSecret
		break
	case AccessTokenType:
		secret = config.AccessTokenSecret
		break
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token")
	}

	return &JWTPayload{
		UserID: uint(claims["sub"].(float64)),
		Email:  claims["email"].(string),
		Role:   claims["aud"].(string),
	}, nil
}
