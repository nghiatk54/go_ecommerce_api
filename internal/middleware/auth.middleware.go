package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/goEcommerceApi/pkg/response"
)

func AuthenMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "valid_token" {
		response.ErrorResponse(c, response.ERR_CODE_INVALID_TOKEN)
		c.Abort()
		return
	}
	c.Next()
}