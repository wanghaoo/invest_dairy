package util

import (
	"fmt"
)

// Qing: 2017-07-24
// 内部错误封装
type InnerError struct {
	Desc   string
	Src    string
	Origin error
}

func (o InnerError) Error() string {
	return fmt.Sprintf("Inner Error: %v [%v]: %v", o.Desc, o.Src, o.Origin)
}

