package exceptin

// ErrorCode 接口定义了获取错误代码和描述的方法。
type ErrorCode interface {
	GetCode() int
	GetDesc() string
}
