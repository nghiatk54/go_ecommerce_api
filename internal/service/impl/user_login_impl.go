package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/constant"
	"github.com/nghiatk54/go_ecommerce_api/internal/database"
	"github.com/nghiatk54/go_ecommerce_api/internal/model"
	"github.com/nghiatk54/go_ecommerce_api/internal/util"
	"github.com/nghiatk54/go_ecommerce_api/internal/util/auth"
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

// *********************************** two factor authentication ***********************************
// check two factor authentication is enabled
func (s *sUserLogin) IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, rs bool, err error) {
	return response.ErrCodeSuccess, true, nil
}

// set up two factor authentication
func (s *sUserLogin) SetUpTwoFactorAuth(ctx context.Context, in *model.SetUpTwoFactorAuthInput) (codeResult int, err error) {
	// 1. check is two factor authentication is enabled
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}
	if isTwoFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthSetupFailed, fmt.Errorf("two factor authentication is already enabled")
	}
	// 2. create new type authentication
	err = s.r.EnabledTwoFactorTypeEmail(ctx, database.EnabledTwoFactorTypeEmailParams{
		UserID:            uint32(in.UserId),
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		TwoFactorEmail:    sql.NullString{String: in.TwoFactorEmail, Valid: true},
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}
	// 3. save otp to redis
	keyUserTwoFactor := crypto.GetHash(fmt.Sprintf("2fa: %d", in.UserId))
	otpNew := random.GenerateSixDigitOtp()
	go func() {
		err = global.Rdb.Set(ctx, keyUserTwoFactor, otpNew, time.Duration(constant.TIME_OTP_REGISTER_MINUTE)*time.Minute).Err()
		if err != nil {
			global.Logger.Error("Failed to set otp to redis", zap.Error(err))
		} else {
			global.Logger.Info("Otp set to redis successfully")
		}
	}()
	// 4. send otp to email
	err = send_to.SendTextEmailOtp([]string{in.TwoFactorEmail}, constant.HOST_EMAIL, strconv.Itoa(otpNew))
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}
	return response.ErrCodeSuccess, nil
}

// verify two factor authentication
func (s *sUserLogin) VerifyTwoFactorAuth(ctx context.Context, in *model.VerifyTwoFactorAuthInput) (codeResult int, err error) {
	// 1. check is two factor authentication is enabled
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	if isTwoFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("two factor authentication is enabled")
	}
	// 2. check otp in redis available
	keyUserTwoFactor := crypto.GetHash(fmt.Sprintf("2fa: %d", in.UserId))
	otpVerifyAuth, err := global.Rdb.Get(ctx, keyUserTwoFactor).Result()
	if err == redis.Nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("key %s not exists", keyUserTwoFactor)
	} else if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	// 3. check otp is correct
	if otpVerifyAuth != in.TwoFactorCode {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("otp does not match")
	}
	// 4. update status is verified
	err = s.r.UpdateTwoFactorStatus(ctx, database.UpdateTwoFactorStatusParams{
		UserID:            uint32(in.UserId),
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	// 5. remove otp from redis
	_, err = global.Rdb.Del(ctx, keyUserTwoFactor).Result()
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	return response.ErrCodeSuccess, nil
}

// *********************************** end two factor authentication ***********************************

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

// login
func (s *sUserLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out *model.LoginOutput, err error) {
	out = &model.LoginOutput{}
	// 1. check user is exists in user base
	userBase, err := s.r.GetOneUserInfo(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// 2. check user password is correct
	if !crypto.MatchingPassword(userBase.UserPassword, in.UserPassword, userBase.UserSalt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match password")
	}
	// 3. check two factor authentication
	isTwoFactorEnabled, err := s.r.IsTwoFactorEnabled(ctx, uint32(userBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("get two factor authentication failed")
	}
	if isTwoFactorEnabled > 0 {
		// save otp to redis
		keyUserLoginTwoFactor := crypto.GetHash(fmt.Sprintf("2fa:otp:%d", userBase.UserID))
		otpNew := random.GenerateSixDigitOtp()
		err = global.Rdb.Set(ctx, keyUserLoginTwoFactor, otpNew, time.Duration(constant.TIME_OTP_REGISTER_MINUTE)*time.Minute).Err()
		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("set otp to redis failed")
		}
		// send otp to email
		infoUserTwoFactor, err := s.r.GetTwoFactorMethodByUserIdAndAuthType(ctx, database.GetTwoFactorMethodByUserIdAndAuthTypeParams{
			UserID:            uint32(userBase.UserID),
			TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		})
		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("get two factor method by user id and auth type failed")
		}
		go func() {
			err = send_to.SendTextEmailOtp([]string{infoUserTwoFactor.TwoFactorEmail.String}, constant.HOST_EMAIL, strconv.Itoa(otpNew))
			if err != nil {
				global.Logger.Error("Failed to send otp to email", zap.Error(err))
			} else {
				global.Logger.Info("Otp sent to email successfully")
			}
		}()
		out.Message = fmt.Sprintf("send OTP 2FA to email: %s, please check your email", infoUserTwoFactor.TwoFactorEmail.String)
		return response.ErrCodeSuccess, out, nil
	}
	// 4. update user login time
	go func() {
		err := s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
			UserLoginIp: sql.NullString{String: "127.0.0.1", Valid: true},
			UserAccount: in.UserAccount,
		})
		if err != nil {
			global.Logger.Error("Failed to update user login time", zap.Error(err))
		} else {
			global.Logger.Info("User login time updated successfully")
		}
	}()
	// 5. create uuid user
	subToken := util.GenerateCliTokenUuid(int(userBase.UserID))
	global.Logger.Info("subToken", zap.String("subToken", subToken))
	// 6. get user info table
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// 7. convert user info to json
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("convert user info to json failed %v", err)
	}
	// 8. save user info to redis with key = subToken
	err = global.Rdb.Set(ctx, subToken, infoUserJson, time.Duration(constant.TIME_2FA_VERIFY_HOUR)*time.Hour).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// 9. create token
	out.Token, err = auth.CreateToken(subToken)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	return response.ErrCodeSuccess, out, nil
}
