package util

import (
  "fmt"
  "math/rand"
  "strconv"
  "time"
)

func Uint64PtrToInt(input *uint64) int {
	if input != nil {
		return int(*input)
	}
	return 0
}

func RandRange(min int, max int) int {
  return rand.Intn(max - min) + min
}

// 运算保留2位小数
func Decimal2(value float64) float64 {
  // return math.Trunc(value*1e2+0.5) * 1e-2
  value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
  return value
}

func Decimal2String(value float64) string {
  return fmt.Sprintf("%.2f", value)
}

func RandInt64(min, max int64) int64 {
  if min >= max || min == 0 || max == 0 {
    return max
  }
  return rand.Int63n(max-min) + min
}

func TimeFromSeconds(seconds int64) time.Time {
  var zero time.Time
  if seconds == 0 {
    return zero
  }
  return time.Unix(0, seconds*int64(time.Second))
}

func GetAge(birthDay string) int {
  birth, err := time.Parse("02/01/2006", birthDay)
  if err != nil {
    fmt.Println("time parse age error: ", err.Error())
    return 0
  }
  now := time.Now()
  age := now.Year() - birth.Year()
  if now.Month() < birth.Month() {
    age -= 1
  }
  return age
}