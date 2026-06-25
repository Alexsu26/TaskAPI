package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenManager struct {
	secret     string
	expiration time.Duration
}

var (
	ErrTokenInvalid = errors.New("auth failed, token is invalid")
	ErrTokenExpired = errors.New("auth failed, token is expired")
)

func NewTokenManager(secret string, expiration time.Duration) *TokenManager {
	return &TokenManager{
		secret:     secret,
		expiration: expiration,
	}
}

func (m *TokenManager) GenerateToken(userID int64) (string, error) {
	if userID <= 0 {
		return "", ErrTokenInvalid
	}
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.expiration)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                   // 签发时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secret))
}

func (m *TokenManager) ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, ErrTokenInvalid
		}
		return []byte(m.secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}
	if !token.Valid || claims.UserID <= 0 {
		return nil, ErrTokenInvalid
	}
	return claims, nil
}
