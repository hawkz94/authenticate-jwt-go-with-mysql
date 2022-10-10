package middlewares

import (
	"tecmentor-api/services"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Baerer_schema = "Baerer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Baerer_schema):]

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}

	}
}
