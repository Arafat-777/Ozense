package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

)

var jwtKey = []byte("your_secret_key")

func AuthMiddleware(roleRequired string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не предоставлен"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			if roleRequired != "" {
				role, ok := claims["Role"].(string)
				if !ok || role != roleRequired {
					c.JSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав"})
					c.Abort()
					return
				}
			}

			c.Set("email", claims["Email"])
			c.Set("role", claims["Role"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		c.Next()
	}
}
