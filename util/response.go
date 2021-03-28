package util

type Response struct {
	Errno int         `json:"errno"`          // 必需 错误码。正常返回0 异常返回560 错误提示561对应errorInfo
	Data  interface{} `json:"data,string"`    // 必需 返回数据内容。 如果有返回数据，可以是字符串或者数组JSON等等
	Page  *Pagination `json:"page,omitempty"` // 非必需 分页信息
}

func ResultData(data interface{}, err error) *Response {
	var result = &Response{}
	if err != nil {
		result.Errno = 501
		result.Data = err.Error()
		return result
	}
	if data == nil {
		data = "success"
	}
	result.Data = data
	return result
}
