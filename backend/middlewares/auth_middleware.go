package middlewares

import (
	"backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(utils.GetEnv("JWT_SECRET", "fallbacksecret"))

type Claims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}

// ใช้ตรวจสอบ JWT Token ก่อนให้เข้า API
func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		
		// ตรวจสอบ Authorization Header
		authHeader := c.GetHeader("Authorization")
		
		// ต้องมีค่า และต้องขึ้นต้นด้วย "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort() // หยุด
			return 
		}

		// ดึง token ออกมา
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// parse และตรวจ
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func (token *jwt.Token) (interface{}, error) {
			// คืนค่า secret key
			return jwtSecret, nil
		})

		// ตรวจ token เสีย / หมดอายุ / ไม่ถูกต้อง
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return 
		}

		// ดึง Claims ออกมา
		if claims, ok := token.Claims.(*Claims); ok {
			// เก็บ userId ไว้ใน Gin Context
			c.Set("userId", claims.UserID)
			c.Next()
			return 
		} 

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
	}
}