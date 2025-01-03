package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Payload
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: key must be at least %d long", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         payload.ID,
		"username":     payload.Username,
		"issued_at":  payload.IssuedAt,
		"expired_at": payload.ExpiredAt,
	})

	return jwtToken.SignedString([]byte(maker.secretKey))
}


func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	claims := &CustomClaims{}

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims.HasExpired() {
		return nil, ErrExpiredToken
	}

	if !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	return &claims.Payload, nil
}
