package response

// http status code
const (
	ERR_CODE_SUCCESS = 20001 // Success
	ERR_CODE_PARAM_INVALID = 20003 // Email is invalid
	ERR_CODE_INVALID_TOKEN = 30001 // Invalid token
)

// message
var msg = map[int]string{
	ERR_CODE_SUCCESS: "success",
	ERR_CODE_PARAM_INVALID: "Email is invalid",
	ERR_CODE_INVALID_TOKEN: "Invalid token",
}