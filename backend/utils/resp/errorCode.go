package resp

// 定义错误枚举
type ErrorCode int

const (
	Success       ErrorCode = 200
	Error         ErrorCode = 201
	InvalidInput  ErrorCode = 400
	Unauthorized  ErrorCode = 401
	Forbidden     ErrorCode = 403
	NotFound      ErrorCode = 404
	InternalError ErrorCode = 500
)

// 错误码和错误信息的映射
var errorMessages = map[ErrorCode]string{
	Success:       "ok",
	Error:         "error",
	InvalidInput:  "Invalid input",
	Unauthorized:  "Unauthorized",
	Forbidden:     "Forbidden",
	NotFound:      "Not Found",
	InternalError: "Internal Server Error",
}

// Message 获取错误信息
func (e ErrorCode) Message() string {
	if msg, exists := errorMessages[e]; exists {
		return msg
	}
	return "Unknown error"
}
