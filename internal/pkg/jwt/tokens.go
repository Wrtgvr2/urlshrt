package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

var AccessTokenLifetime = time.Minute * 15
var RefreshTokenLifetime = time.Hour * 24 * 30
var signingMethod = jwt.SigningMethodHS256
var secretKey = os.Getenv("JWT_SECRET")

func createAccessToken(userID uint64) (string, error) {
	tokenClaims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenLifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(signingMethod, tokenClaims)
	tokenStr, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func createRefreshToken(userID uint64) (string, error) {
	tokenClaims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTokenLifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        uuid.New().String(),
		},
	}
	refreshToken := jwt.NewWithClaims(signingMethod, tokenClaims)
	tokenStr, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// Return: access token, refresh token, error
func CreateTokens(userID uint64) (string, string, error) {
	accessToken, err := createAccessToken(userID)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := createRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}
