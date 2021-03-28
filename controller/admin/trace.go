package admin

import (
	"bytes"
	"invest_dairy/common"
	"invest_dairy/service/admin"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func traceOp(operationDesc string) func(user *admin.UserInfoVo, ctx echo.Context) {
	return func(user *admin.UserInfoVo, ctx echo.Context) {
		// 取出body内容
		var save, read, _ = drainBody(ctx.Request().Body)
		readBytes, _ := ioutil.ReadAll(read)
		ctx.Request().Body = save
		common.Mlog.WithField("username", user.NickName).
			WithField("userId", strconv.Itoa(user.Id)).
			WithField("operationDesc", operationDesc).
			WithField("url", ctx.Request().URL.String()).
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
