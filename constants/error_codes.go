package constants

// ErrorCode 定义统一错误码结构
type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 常见 Web 请求错误码
var (
	ErrOK                 = ErrorCode{Code: 200, Message: "请求成功"}
	ErrBadRequest         = ErrorCode{Code: 400, Message: "错误的请求"}
	ErrUnauthorized       = ErrorCode{Code: 401, Message: "未授权访问"}
	ErrForbidden          = ErrorCode{Code: 403, Message: "禁止访问"}
	ErrNotFound           = ErrorCode{Code: 404, Message: "资源未找到"}
	ErrInternalServer     = ErrorCode{Code: 500, Message: "服务器内部错误"}
	ErrServiceUnavailable = ErrorCode{Code: 503, Message: "服务不可用"}
)
