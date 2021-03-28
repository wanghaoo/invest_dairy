package util

func IsNoRowsError(err error) bool {
  if err == nil {
    return false
  }
	return "sql: no rows in result set" == err.Error()
}
