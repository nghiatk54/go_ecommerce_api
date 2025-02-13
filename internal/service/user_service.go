package service

import (
	"context"

	"github.com/nghiatk54/go_ecommerce_api/internal/model"
)

// define interface
type (
	IUserLogin interface {
		Login(ctx context.Context) error
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOtp(ctx context.Context, in *model.VerifyOtpInput) (out *model.VerifyOtpOutput, err error)
		UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error)
	}
	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}
	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

// define variable for interface
var (
	localUserAdmin IUserAdmin
	localUserInfo  IUserInfo
	localUserLogin IUserLogin
)

// init and get UserAdmin
func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement localUserAdmin not found for interface IUserAdmin")
	}
	return localUserAdmin
}
func InitUserAdmin(userAdmin IUserAdmin) {
	localUserAdmin = userAdmin
}

// init and get UserInfo
func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement localUserInfo not found for interface IUserInfo")
	}
	return localUserInfo
}
func InitUserInfo(userInfo IUserInfo) {
	localUserInfo = userInfo
}

// init and get UserLogin
func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement localUserLogin not found for interface IUserLogin")
	}
	return localUserLogin
}
func InitUserLogin(userLogin IUserLogin) {
	localUserLogin = userLogin
}
