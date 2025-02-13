package impl

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/constant"
	"github.com/nghiatk54/go_ecommerce_api/internal/database"
	"github.com/nghiatk54/go_ecommerce_api/internal/model"
	"github.com/nghiatk54/go_ecommerce_api/internal/util"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/crypto"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/random"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/send_to"
	"github.com/nghiatk54/go_ecommerce_api/pkg/response"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// define struct user login
type sUserLogin struct {
	r *database.Queries
}

// create new user login
func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// register
func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	// 1. hash email or mobile phone
	fmt.Printf("VerifyKey: %s\n", in.VerifyKey)
	fmt.Printf("VerifyType: %d\n", in.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)
	// 2. check user exists in user base
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}
	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}
	// 3. check Otp exists
	userKey := util.GetUserKey(hashKey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()
	// util ...
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get failed:", err)
		return response.ErrInvalidOtp, err
	case otpFound != "":
		return response.ErrCodeOtpNotExists, fmt.Errorf("otp exists! but not registered")
	}
	// 4. generate Otp
	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("Otp is::%d\n", otpNew)
	// 5. save Otp to redis with expired time
	err = global.Rdb.Set(ctx, userKey, strconv.Itoa(otpNew), time.Duration(constant.TIME_OTP_REGISTER_MINUTE)*time.Minute).Err()
	if err != nil {
		return response.ErrInvalidOtp, err
	}
	// 6. send otp
	switch in.VerifyType {
	case constant.EMAIL:
		err = send_to.SendTextEmailOtp([]string{in.VerifyKey}, constant.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		// 7. save Otp to MySQL
		result, err := s.r.InsertOtpVerify(ctx, database.InsertOtpVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashKey,
		})
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		// 8. get last id
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		global.Logger.Info("Insert otp verify success", zap.Int64("Last id verify user: ", lastIdVerifyUser))
		return response.ErrCodeSuccess, nil
	case constant.MOBILE_PHONE:
		return response.ErrCodeSuccess, nil
	}
	return response.ErrCodeSuccess, nil
}

// login
func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

// verify otp
func (s *sUserLogin) VerifyOtp(ctx context.Context, in *model.VerifyOtpInput) (out *model.VerifyOtpOutput, err error) {
	// 0. set output
	out = &model.VerifyOtpOutput{}
	// 1. hash key
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	// 2. check otp is correct in redis
	otpFound, err := global.Rdb.Get(ctx, util.GetUserKey(hashKey)).Result()
	if err != nil {
		return out, err
	}
	if in.VerifyCode != otpFound {
		// if otp is not match in 1 minute?
		return out, fmt.Errorf("otp is not match")
	}
	// 3. check user is exists in user verify
	infoOtp, err := s.r.GetInfoOtp(ctx, hashKey)
	if err != nil {
		return out, err
	}
	// 4. update status is verified
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}
	// 5. return success
	out.Token = infoOtp.VerifyKeyHash
	out.Message = "Verify success"
	return out, nil
}

// update password register
func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error) {
	// 1. check token is exists in database
	infoOtp, err := s.r.GetInfoOtp(ctx, token)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	// 2. check is verified
	if infoOtp.IsVerified.Int32 == 0 {
		return response.ErrCodeUserOtpNotExists, fmt.Errorf("user Otp not verified")
	}
	// 3. update user base
	userBase := database.AddUserBaseParams{}
	userBase.UserAccount = infoOtp.VerifyKey
	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	userBase.UserSalt = userSalt
	userBase.UserPassword = crypto.HashPassword(password, userSalt)
	// 4. insert user base to table in database
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	user_id, err := newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	// 5. add user info have user id
	newUserInfo, err := s.r.AddUserHaveUserId(ctx, database.AddUserHaveUserIdParams{
		UserID:               uint64(user_id),
		UserAccount:          infoOtp.VerifyKey,
		UserNickname:         sql.NullString{String: infoOtp.VerifyKey, Valid: true},
		UserAvatar:           sql.NullString{String: "", Valid: true},
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOtp.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	})
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	lastIdUserInfo, err := newUserInfo.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	return int(lastIdUserInfo), nil
}
