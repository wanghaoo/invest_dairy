package controller

import (
	"bytes"
	"invest_dairy/common"
	"invest_dairy/service"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func traceOp(operationDesc string) func(user *service.UserInfoVo, ctx echo.Context) {
	return func(user *service.UserInfoVo, ctx echo.Context) {
		// 取出body内容
		var save, read, _ = drainBody(ctx.Request().Body)
		readBytes, _ := ioutil.ReadAll(read)
		ctx.Request().Body = save
		logs := common.Mlog.WithField("operationDesc", operationDesc)
		if user != nil {
			logs = logs.WithField("username", user.Phone).
				WithField("userId", strconv.Itoa(user.Id))
		}
		logs.WithField("url", ctx.Request().URL.String()).
			WithField("body", string(readBytes)).
			Infof("")
	}
}

func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
