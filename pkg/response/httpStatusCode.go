package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001
	ErrInvalidOtp       = 30002
	ErrSendEmailOtp     = 30003
	// Register code
	ErrCodeUserHasExists = 50001
	// login code
	ErrCodeOtpNotExists     = 60009
	ErrCodeUserOtpNotExists = 60008
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid!",
	ErrInvalidToken:     "Invalid token!",
	ErrInvalidOtp:       "OTP is invalid!",
	ErrSendEmailOtp:     "Send email Otp failed!",
	// Register message
	ErrCodeUserHasExists: "User already exists!",
	// login message
	ErrCodeOtpNotExists:     "Otp exists! but not registered!",
	ErrCodeUserOtpNotExists: "User Otp not exists!",
}
