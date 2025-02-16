package account

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/model"
	"github.com/nghiatk54/go_ecommerce_api/internal/service"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/context"
	"github.com/nghiatk54/go_ecommerce_api/pkg/response"
	"go.uber.org/zap"
)

var TwoFa = new(cUser2Fa)

type cUser2Fa struct{}

// User setup two factor authentication documentation
// @Summary      User setup two factor authentication
// @Description  User setup two factor authentication
// @Tags         account 2Fa
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization token"
// @Param        payload body model.SetUpTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two_factor/setup [post]
func (c *cUser2Fa) SetUpTwoFactorAuth(ctx *gin.Context) {
	var params model.SetUpTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "Missing or invalid set up two factor authentication parameters")
		return
	}
	// get user id from uuid(token)
	userId, err := context.GetUserIdFromUuid(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Get user id from uuid failed")
		return
	}
	global.Logger.Info("userId: ", zap.Uint64("userId", userId))
	params.UserId = uint32(userId)
	codeResult, err := service.UserLogin().SetUpTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, codeResult, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, nil)
}

// User verify two factor authentication documentation
// @Summary      User verify two factor authentication
// @Description  User verify two factor authentication
// @Tags         account 2Fa
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization token"
// @Param        payload body model.VerifyTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two_factor/verify [post]
func (c *cUser2Fa) VerifyTwoFactorAuth(ctx *gin.Context) {
	var params model.VerifyTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "Missing or invalid verify two factor authentication parameters")
		return
	}
	// get user id from uuid(token)
	userId, err := context.GetUserIdFromUuid(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed, "Get user id from uuid failed")
		return
	}
	global.Logger.Info("userId: ", zap.Uint64("userId", userId))
	params.UserId = uint32(userId)
	codeResult, err := service.UserLogin().VerifyTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, codeResult, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, nil)
}
