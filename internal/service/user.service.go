package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/repo"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/crypto"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/random"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/send_to"
	"github.com/nghiatk54/go_ecommerce_api/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepo
	userAuthRepo repo.IUserAuthRepo
}

// Register user
func (us *userService) Register(email string, purpose string) int {
	// 0. hash email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("Hash email is ::%s", hashEmail)

	// 5. check OTP is available

	// 6. check user spam

	// 1. check email has exists in database
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2. create new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("OTP is ::%d\n", otp)

	// 3. save OTP to redis with expiration time
	err := us.userAuthRepo.AddOtp(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOtp
	}
	// 4. send OTP to email
	err = send_to.SendTemplateEmailOtp([]string{email}, global.Config.Smtp.From, "otp_auth.html", map[string]interface{}{
		"otp": strconv.Itoa(otp),
	})
	if err != nil {
		return response.ErrSendEmailOtp
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepo, userAuthRepo repo.IUserAuthRepo) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}
