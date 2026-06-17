package helper

import (
    "os"
    "time"
    "errors"
    "github.com/golang-jwt/jwt/v5"
)

var ErrTokenInvalid = errors.New("invalid token") 
var ErrTokenExpired = errors.New("token expired")

type Claims struct {
    UserID uint   `json:"user_id"`
    Phone  string `json:"phone"`
    jwt.RegisteredClaims
}

func GenerateToken(userID uint, phone string) (string, error) {
    claims := Claims{
        UserID: userID,
        Phone:  phone,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(tokenStr string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, ErrTokenInvalid
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        return nil, ErrTokenInvalid
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}