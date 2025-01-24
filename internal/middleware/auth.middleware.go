package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go_ecommerce_api/pkg/response"
)

func AuthenMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "valid_token" {
		response.ErrorResponse(c, response.ErrInvalidToken)
		c.Abort()
		return
	}
	c.Next()
}
