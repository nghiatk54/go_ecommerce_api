package account

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/model"
	"github.com/nghiatk54/go_ecommerce_api/internal/service"
	"github.com/nghiatk54/go_ecommerce_api/pkg/response"
	"go.uber.org/zap"
)

// management controller login user
var Login = new(cUserLogin)

type cUserLogin struct{}

// User login documentation
// @Summary      User login
// @Description  User login
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.LoginInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/login [post]
func (c *cUserLogin) Login(ctx *gin.Context) {
	var params model.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeResult, dataResult, err := service.UserLogin().Login(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, codeResult, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, dataResult)
}

// User registration documentation
// @Summary      User registration
// @Description  When user is registered send otp to email or mobile phone
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/register [post]
func (c *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Register error user Otp", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeStatus, nil)
}

// User verify otp documentation
// @Summary      Verify otp login by user
// @Description  Verify otp login by user
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.VerifyOtpInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/verify_account [post]
func (c *cUserLogin) VerifyOtp(ctx *gin.Context) {
	var params model.VerifyOtpInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	result, err := service.UserLogin().VerifyOtp(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidOtp, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}

// User update password when register documentation
// @Summary      User update password when register
// @Description  User update password when register
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdatePasswordRegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/update_password_register [post]
func (c *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params model.UpdatePasswordRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	result, err := service.UserLogin().UpdatePasswordRegister(ctx, params.UserToken, params.UserPassword)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
