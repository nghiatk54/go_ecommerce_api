package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/constant"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/auth"
	"github.com/nghiatk54/go_ecommerce_api/pkg/response"
	"go.uber.org/zap"
)

func AuthenMiddleware(c *gin.Context) {
	// get the request url
	uri := c.Request.URL.Path
	global.Logger.Info("uri request: ", zap.String("uri", uri))
	// get jwt token from header
	jwtToken, ok := auth.ExtractBearerToken(c)
	if !ok {
		response.ErrorResponseAbort(c, response.ErrCodeUnauthorized)
		return
	}
	// validate jwt token by subject
	claims, err := auth.VerifyTokenSubject(jwtToken)
	if err != nil {
		response.ErrorResponseAbort(c, response.ErrCodeUnauthorized)
		return
	}
	// update claims to context
	global.Logger.Info("claims::UUID: ", zap.String("uuid", claims.Subject))
	ctx := context.WithValue(c.Request.Context(), constant.SUBJECT_UUID, claims.Subject)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
