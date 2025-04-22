package jwt

import (
	"fmt"
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

func createAccessToken(userID uint64, secretKey []byte) (string, error) {
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

func createRefreshToken(userID uint64, secretKey []byte) (string, error) {
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
	var secretKey = []byte(os.Getenv("JWT_SECRET"))
	fmt.Printf("JWT_SECRET from env:%ssecretKey from variable: %s\n", os.Getenv("JWT_SECRET"), secretKey)
	accessToken, err := createAccessToken(userID, secretKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := createRefreshToken(userID, secretKey)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
