package common

type ResponseData struct {
  Errno int         `json:"errno"`       // 必需 错误码。正常返回0 异常返回560 错误提示561对应errorInfo
  Data  interface{} `json:"data,string"` // 可选 返回数据内容。 如果有返回数据，可以是字符串或者数组JSON等等
}

type Page struct {
  PageNo   int `json:"pageNo"`
  PageSize int `json:"pageSize"`
  Total    int    `json:"total"`
}

type PageData struct {
  Rows interface{} `json:"rows"`
  Page Page        `json:"page"`
}

func SetResult(data interface{}, pageNo int, pageSize int, total int) *ResponseData {
  if pageNo == 0 {
    pageNo = 1
  }
  if pageSize == 0 {
    pageSize = 10
  }
  return SetData(PageData{Rows: data, Page: Page{PageNo: pageNo, PageSize: pageSize, Total: total}})
}

func SetError(message string) *ResponseData {
  response := new(ResponseData)
  response.Errno = 501
  response.Data = message
  return response
}
func SetErrorNo(errorNo int, message string) *ResponseData {
  response := new(ResponseData)
  response.Errno = errorNo
  response.Data = message
  return response
}

func CommonError() *ResponseData {
  response := new(ResponseData)
  response.Errno = 501
  response.Data = "Please try again later"
  return response
}

func SetData(data interface{}) *ResponseData {
  response := new(ResponseData)
  response.Errno = 200
  response.Data = data
  return response
}

func CommonSuccess() *ResponseData {
  response := new(ResponseData)
  response.Errno = 200
  response.Data = "Operation successful"
  return response
}
