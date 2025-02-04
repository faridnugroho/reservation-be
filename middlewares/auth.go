package middlewares

import (
	"net/http"
	"reservation/dto"
	"reservation/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Status:  http.StatusUnauthorized,
				Message: "No JWT token provided",
			})
			c.Abort()
			return
		}

		// Memisahkan token dari "Bearer <token>"
		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Status:  http.StatusUnauthorized,
				Message: "Invalid token format",
			})
			c.Abort()
			return
		}

		claims, err := jwt.DecodeToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Status:  http.StatusUnauthorized,
				Message: "Failed to decode JWT token",
			})
			c.Abort()
			return
		}

		// Menyimpan data user ke dalam context
		c.Set("currentUser", claims)

		// Lanjut ke handler berikutnya
		c.Next()
	}
}
