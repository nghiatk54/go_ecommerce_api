package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrInvalidToken     = 30001
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid!",
	ErrInvalidToken:     "Invalid token!",
}
