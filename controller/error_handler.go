package controller

import (
	"invest_dairy/bizerrors"
	"invest_dairy/common"
	"invest_dairy/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JsonError struct {
	httpStatus int
	json       []byte
	message    string
}

func (e *JsonError) Error() string {
	return e.message
}

func (e *JsonError) Json() []byte {
	return e.json
}

func (e *JsonError) HttpStatus() int {
	return e.httpStatus
}

// BindRoutes func
func init() {

	util.EchoInst.HTTPErrorHandler = func(err error, c echo.Context) {
		httpStatus := http.StatusInternalServerError
		common.Mlog.WithField("error_handler", strconv.Itoa(httpStatus)).Errorf("Method: %v \nURL: %v \nError: %v", c.Request().Method, c.Request().URL, err)
		if c.Response().Committed {
			return
		}

		if bizError, ok := err.(*bizerrors.BizError); ok {
			c.JSON(bizError.HttpStatus(), map[string]interface{}{
				"errno": bizError.Code(),
				"msg":   bizError.Error(),
			})
			return
		}
		c.JSON(httpStatus, map[string]interface{}{
			"errno": httpStatus,
			"msg":   err.Error(),
		})
	}
}
