package jwt

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type tokenClaims struct {
	UserID string `json:"sub"`
	jwt.StandardClaims
}

var AccessTokenLifetime = time.Minute * 15
var RefreshTokenLifetime = time.Hour * 24 * 30
var signingMethod = jwt.SigningMethodHS256

func createAccessToken(userID uint64, secretKey []byte) (string, error) {
	tokenClaims := tokenClaims{
		UserID: fmt.Sprint(userID),
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
	tokenClaims := tokenClaims{
		UserID: fmt.Sprint(userID),
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

func keyFunc(t *jwt.Token) (any, error) {
	var secretKey = []byte(os.Getenv("JWT_SECRET"))

	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}

	return secretKey, nil
}

func ParseTokenStr(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, keyFunc)

	return token, err
}

func ValidateToken(tokenStr string) error {
	token, err := ParseTokenStr(tokenStr)
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func GetTokenClaims(tokenStr string) (*tokenClaims, error) {
	claims := &tokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, keyFunc)
	if !token.Valid {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func GetUserIdFromToken(tokenStr string) (uint64, error) {
	claims, err := GetTokenClaims(tokenStr)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseUint(claims.UserID, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetTokenExpirationUnixTime(tokenStr string) (int64, error) {
	claims, err := GetTokenClaims(tokenStr)
	if err != nil {
		return 0, err
	}

	return claims.ExpiresAt, nil
}

func GetJtiFromRefreshToken(tokenStr string) (uuid.UUID, error) {
	claims, err := GetTokenClaims(tokenStr)
	if err != nil {
		return uuid.Nil, err
	}
	jti, err := uuid.Parse(claims.Id)
	if err != nil {
		return uuid.Nil, err
	}
	return jti, nil
}
