package utils

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(secret string, expireAfter time.Duration, userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.FormatUint(uint64(userID), 10),
		"exp": jwt.NewNumericDate(time.Now().Add(expireAfter)),
	})

	return token.SignedString([]byte(secret))
}

func ParseToken(secret string, tokenString string) (uint, time.Time, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return 0, time.Time{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userIdStr, err := claims.GetSubject()
		if err != nil {
			return 0, time.Time{}, err
		}
		userID, err := strconv.ParseUint(userIdStr, 10, 64)
		if err != nil {
			return 0, time.Time{}, err
		}
		expireAt, err := claims.GetExpirationTime()
		if err != nil {
			return 0, time.Time{}, err
		}
		return uint(userID), expireAt.Time, err
	}
	return 0, time.Time{}, fmt.Errorf("Unexpected token")
}
