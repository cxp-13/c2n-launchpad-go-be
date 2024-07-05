package exceptin

import (
	"fmt"
)

// CommonException 是一个自定义错误类型，它包装了 ErrorCode。
type CommonException struct {
	errorCode ErrorCode
}

// NewCommonException 返回一个新的 CommonException 实例。
func NewCommonException(errorCode ErrorCode) *CommonException {
	return &CommonException{
		errorCode: errorCode,
	}
}

// Error 实现了 error 接口，返回错误描述。
func (e *CommonException) Error() string {
	return fmt.Sprintf("error %d: %s", e.errorCode.GetCode(), e.errorCode.GetDesc())
}

// GetErrorCode 返回错误代码。
func (e *CommonException) GetErrorCode() ErrorCode {
	return e.errorCode
}
