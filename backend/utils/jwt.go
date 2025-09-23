package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims ใช้เป็น payload ใน JWT
type Claims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}

// GenerateToken สำหรับ userId
func GenerateToken(userId uint) (string, error) {
	// โหลด secret จาก env
	secret := GetEnv("JWT_SECRET", "fallbacksecret")

	// TTL (อายุ token)
	ttlStr := GetEnv("JWT_TTL", "3600")
	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		ttl = 3600
	}

	// เวลาหมดอายุ
	expirationTime := time.Now().Add(time.Duration(ttl) * time.Second)

	// claims
	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:	jwt.NewNumericDate(time.Now()),
		},
	}

	// สร้าง token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// เซ็นด้วย secret
	return token.SignedString([]byte(secret))
}

// ParseToken ใช้ตรวจและดึง claims จาก JWT
func ParseToken(tokenStr string) (*Claims, error) {
	secret := GetEnv("JWT_SECRET", "fallbacksecret")

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}