package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go_ecommerce_api/internal/service"
	"github.com/nghiatk54/go_ecommerce_api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	fmt.Println("My handler")
	response.SuccessResponse(c, 20001, []string{"cs7", "m10", "nghiatk54"})
	// response.ErrorResponse(c, 20003)
}
