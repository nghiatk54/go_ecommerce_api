package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponseData struct {
	Code   int         `json:"code"`
	Error  string      `json:"error"`
	Detail interface{} `json:"detail"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, messages ...string) {
	message := msg[code]
	if len(messages) > 0 {
		message = messages[0]
	}
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func ErrorResponseAbort(c *gin.Context, code int, messages ...string) {
	message := msg[code]
	if len(messages) > 0 {
		message = messages[0]
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
