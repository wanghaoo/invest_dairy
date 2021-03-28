package common

import "testing"

func Test_log(t *testing.T) {
  Mlog.Debug("哈哈哈哈")
  Mlog.Error("哈哈哈哈")
  Mlog.Errorf("11111%s","哈哈哈哈")
}
