package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001
	ErrInvalidOtp       = 30002
	ErrSendEmailOtp     = 30003
	// user authentication
	ErrCodeUnauthorized = 40001
	ErrCodeAuthFailed   = 40005
	// Register code
	ErrCodeUserHasExists = 50001
	// login code
	ErrCodeOtpNotExists     = 60009
	ErrCodeUserOtpNotExists = 60008
	// two factor authentication
	ErrCodeTwoFactorAuthSetupFailed  = 80001
	ErrCodeTwoFactorAuthVerifyFailed = 80002
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid!",
	ErrInvalidToken:     "Invalid token!",
	ErrInvalidOtp:       "OTP is invalid!",
	ErrSendEmailOtp:     "Send email Otp failed!",
	// user authentication
	ErrCodeUnauthorized: "Unauthorized!",
	ErrCodeAuthFailed:   "Authentication failed!",
	// Register message
	ErrCodeUserHasExists: "User already exists!",
	// login message
	ErrCodeOtpNotExists:     "Otp exists! but not registered!",
	ErrCodeUserOtpNotExists: "User Otp not exists!",
	// two factor authentication
	ErrCodeTwoFactorAuthSetupFailed:  "Two factor authentication setup failed!",
	ErrCodeTwoFactorAuthVerifyFailed: "Two factor authentication verify failed!",
}
