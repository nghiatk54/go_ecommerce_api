package repo

import (
	"fmt"
	"time"

	"github.com/nghiatk54/go_ecommerce_api/global"
)

type IUserAuthRepo interface {
	AddOtp(email string, otp int, expirationTime int64) error
}

type userAuthRepo struct{}

func (u *userAuthRepo) AddOtp(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepo() IUserAuthRepo {
	return &userAuthRepo{}
}
