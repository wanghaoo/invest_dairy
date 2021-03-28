package bizerrors

import "net/http"

// BizError struct
type BizError struct {
  httpStatus int
  code       string
  message    string
}

// HttpStatus returns http status.
func (e *BizError) HttpStatus() int {
  return e.httpStatus
}

// Code returns code.
func (e *BizError) Code() string {
  return e.code
}

// Error returns message.
func (e *BizError) Error() string {
  return e.message
}

func CustomBizError(message string) *BizError {
  return &BizError{http.StatusOK, "10000", message}
}

var (
  VerifyTokenError = &BizError{http.StatusOK, "505", "Please login"}
)
