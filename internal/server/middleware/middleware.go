package middleware

import (
	"br-lesson-4/internal/server/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AuthMiddleware(jwtSigner auth.HS256Signer) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User Not Authorized",
			})
			c.Abort()
			return
		}

		claims, err := jwtSigner.ParseAccessToken(authHeader, auth.ParseOptions{
			ExpectedIssuer:   jwtSigner.Issuer,
			ExpectedAudience: jwtSigner.Audience,
			AllowedMethods:   []string{"HS256"},
			Leeway:           60 * time.Second,
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
