package utils

import (
	"math"
)

// PageParam 用于定义分页参数的结构体（假设这个结构体已在其他 Go 文件中定义）。
type PageParam struct {
	PageIndex int
	PageSize  int
}

// Page 用于定义泛型分页数据的结构体。
type Page[T any] struct {
	PageIndex int
	PageSize  int
	PageCount int
	DataCount int64
	Data      []T
}

// NewPage 创建并初始化一个新的 Page 实例。
func NewPage(pageParam PageParam, dataCount int64, data []any) (*Page[any], error) {
	page := &Page[any]{
		PageIndex: pageParam.PageIndex,
		PageSize:  pageParam.PageSize,
		DataCount: dataCount,
		PageCount: int(math.Ceil(float64(dataCount) / float64(pageParam.PageSize))),
		Data:      data,
	}

	//if page.PageCount > 0 && page.PageIndex > page.PageCount {
	//	// 如果页码超出总页数，抛出自定义异常
	//	return nil, CommonException{ErrorCode: 0, Message: "Invalid page index."}
	//}

	return page, nil
}
