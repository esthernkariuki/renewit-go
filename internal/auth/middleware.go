package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// =========================
		// GET AUTHORIZATION HEADER
		// =========================

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			c.Abort()
			return
		}

		// =========================
		// CHECK BEARER TOKEN
		// =========================

		parts := strings.Fields(authHeader)

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// =========================
		// PARSE JWT
		// =========================

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				// Make sure the token uses HMAC
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrTokenSignatureInvalid
				}

				return jwtSecret, nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// =========================
		// GET CLAIMS
		// =========================

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			c.Abort()
			return
		}

		// =========================
		// GET USER ID
		// =========================

		userID, ok := claims["user_id"].(float64)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid user ID in token",
			})
			c.Abort()
			return
		}

		// =========================
		// GET ROLE
		// =========================

		role, ok := claims["role"].(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid role in token",
			})
			c.Abort()
			return
		}

		// =========================
		// GET PHONE
		// =========================

		phone, _ := claims["phone"].(string)

		// =========================
		// STORE USER INFORMATION
		// =========================

		c.Set("user_id", userID)
		c.Set("role", role)
		c.Set("phone", phone)

		// Continue request
		c.Next()
	}
}
