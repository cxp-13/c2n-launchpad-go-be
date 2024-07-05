package commonenums

// ReCode 结构体定义了错误代码和描述。
type ReCode struct {
	Code    int    // 状态码
	Message string // 描述信息
}

// GetCode 返回 ReCode 的状态码。
func (r ReCode) GetCode() int {
	return r.Code
}

// GetDesc 返回 ReCode 的描述信息。
func (r ReCode) GetDesc() string {
	return r.Message
}

// 状态常量定义
var (
	SUCCESS               = ReCode{200, "Success"}
	NO_PERMISSION         = ReCode{403, "No Permission"}
	SERVER_ERROR          = ReCode{500, "Internal Server Error"}
	INVALID_PARAMETERS    = ReCode{501, "Invalid Parameters"}
	INVALID_TOKEN         = ReCode{502, "Invalid Token"}
	TOKEN_EXPIRED         = ReCode{502, "The token has expired"}
	FAILED                = ReCode{503, "Failed"}
	DATA_DUPLICATION      = ReCode{505, "Data exists"}
	VERIFICATION_FAILED   = ReCode{506, "Wrong account or password"}
	REQUEST_TOO_FREQUENT  = ReCode{509, "Requests are too frequent"}
	PRIVATE_KEY_EXIST     = ReCode{502, "exist"}
	PRIVATE_KEY_NOT_EXIST = ReCode{503, "key not exist"}
)
