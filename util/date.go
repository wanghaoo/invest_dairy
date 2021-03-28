package util

import "time"

func CovertMinTimeInDay(beginDateTime int64) int64 {
  beginDate := time.Unix(beginDateTime, 0)
  beginDateStr := beginDate.Format("2006-01-02") + " 00:00:00"
  location, _ := time.LoadLocation("Asia/Shanghai")
  beginDate, _ = time.ParseInLocation("2006-01-02 15:04:05", beginDateStr, location)
  return beginDate.Unix()
}

func CovertMaxTimeInDay(endDateTime int64) int64 {
  endDate := time.Unix(endDateTime, 0)
  endDateStr := endDate.Format("2006-01-02") + " 23:59:59"
  location, _ := time.LoadLocation("Asia/Shanghai")
  endDate, _ = time.ParseInLocation("2006-01-02 15:04:05", endDateStr, location)
  return endDate.Unix()
}
