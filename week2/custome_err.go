package week2

import "fmt"

// 错误码
const (
	NotFound = 404001
	Database = 500001
)

// CustomErr 自定义 err
type CustomErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// NewErr 新建一个自定义 err
func NewErr(code int, msg string) *CustomErr {
	return &CustomErr{
		Code: code,
		Msg:  msg,
	}
}

func ErrCode(err error) int {
	if e, ok := err.(*CustomErr); ok {
		return e.Code
	}

	return 0
}

func (ce *CustomErr) Error() string {
	return fmt.Sprintf("[%d] %s", ce.Code, ce.Msg)
}
