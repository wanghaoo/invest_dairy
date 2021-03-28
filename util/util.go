package util

import (
  "math/rand"
  "strings"
  "time"
  "crypto/sha1"
  "fmt"
  "regexp"
  "reflect"
  "strconv"
  "sort"
)

const (
  letterBytes        = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  timeLayout24       = "2006-01-02 15:04:05"
  timeLayoutDate     = "2006-01-02"
)

var sh, _ = time.LoadLocation("Asia/Shanghai")

func CheckStrIsBlank(strParam string) bool {
  return strings.TrimSpace(strParam) == ""
}

func CheckStrPointIsBlank(strPoint *string) bool {
  return strPoint == nil || CheckStrIsBlank(*strPoint)
}

func GetIntValFromBool(boolVal bool) int {
  if boolVal {
    return 1
  }
  return 0
}

func GetBoolValFromInt(intVal int) bool {
  if intVal == 0 {
    return false
  }
  return true
}

func GainRandomString(length int) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = letterBytes[rand.Intn(len(letterBytes))]
  }
  return string(b)
}

func RemoveDuplicates(src []string) []string {
  result := make([]string, 0)

  for i := 0; i < len(src); i++ {
    exists := false
    for v := 0; v < i; v++ {
      if strings.TrimSpace(src[v]) == strings.TrimSpace(src[i]) {
        exists = true
        break
      }
    }
    if !exists {
      result = append(result, strings.TrimSpace(src[i]))
    }
  }
  return result
}

type stringSliceUtil []string

func (p *stringSliceUtil) RemoveDuplicates() *stringSliceUtil {
  *p = RemoveDuplicates(p.StringSlice())
  return p
}

func (o stringSliceUtil) IfContains(testSlice []string) bool {
  return o.IfContainsF(testSlice, func(param1, param2 string) bool {
    return param1 == param2
  })
}

type StringEqualsFunc func(str1, str2 string) bool

func (o stringSliceUtil) IfContainsF(testSlice []string, equalsFunc StringEqualsFunc) bool {
OUTER:
  for _, originElement := range o.StringSlice() {
    containsElement := false
    for _, testElement := range testSlice {
      if equalsFunc(originElement, testElement) {
        containsElement = true
        continue OUTER
      }
    }
    if !containsElement {
      return false
    }
  }
  return true
}

func (o stringSliceUtil) IfContain(testStr string) bool {
  return o.IfContains([]string{testStr})
}

func (o stringSliceUtil) StringSlice() []string {
  return []string(o)
}

/**
 * 生成随机字符串
 * @param  num int
 * @param  kind
    KC_RAND_KIND_NUM   = 0 // 纯数字
    KC_RAND_KIND_LOWER = 1 // 小写字母
    KC_RAND_KIND_UPPER = 2 // 大写字母
    KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
 * @return str string
 */
func GetRandomString(size int, kind int) string {
  ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
  is_all := kind > 2 || kind < 0
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < size; i++ {
    if is_all { // random ikind
      ikind = rand.Intn(3)
    }
    scope, base := kinds[ikind][0], kinds[ikind][1]
    result[i] = uint8(base + rand.Intn(scope))
  }
  return string(result)
}

/*
* 生成sha1
*anthor:vance
*/
func Js_sha1(data string) string {
  h := sha1.New()
  h.Write([]byte(data))
  sha1str1 := h.Sum(nil)
  sha1str2 := fmt.Sprintf("%x", sha1str1)
  return sha1str2

  /*AWSSecretKeyId := "ooxxooxx"
  sha256 := sha256.New
  hash := hmac.New(sha256, []byte(AWSSecretKeyId))
  hash.Write([]byte(data))
  sha := base64.StdEncoding.EncodeToString(hash.Sum(nil))
  sha= url.QueryEscape(sha)
  return sha*/
}

func JoinArticleImages(inputImages []string) string {
  return strings.Join(inputImages, ",")
}

func RandInt(min, max int) int {
  if min >= max || min == 0 || max == 0 {
    return max
  }
  return rand.Intn(max-min) + min
}

func FormatUnixTime(ctime int64) string {
  if ctime == 0 {
    return ""
  }
  tm := time.Unix(ctime, 0)
  return tm.Format("2006-01-02 15:04:05")
}

func FormatUnixHHmm(ctime int64) string {
  tm := time.Unix(ctime, 0)
  return tm.Format("15:04")
}

//判断结构体的所有字段是否都是正整数，目前只判断字段类型是： string, int, ptr
func IsNumber(obj interface{}) bool {
  v := reflect.ValueOf(obj)
  if v.Kind() == reflect.Ptr {
    v = v.Elem()
  }
  count := v.NumField()
  reg := regexp.MustCompile("^[0-9]+$")
  for i := 0; i < count; i++ {
    f := v.Field(i)
    var i string
    switch f.Kind() {
    case reflect.String:
      i = f.String()
      break
    case reflect.Int:
      i = strconv.FormatInt(f.Int(), 10)
      break
    case reflect.Ptr:
      pf := f.Elem().Kind()
      switch pf {
      case reflect.String:
        i = f.Elem().String()
        break
      case reflect.Int:
        i = strconv.FormatInt(f.Elem().Int(), 10)
        break
      }
      break
    default:
      return false
    }
    if len(i) == 0 {
      continue
    }
    if !reg.Match([]byte(i)) {
      return false
    }
  }
  return true
}

func SHNow() time.Time {
  return time.Now().In(sh)
}

func SHFromUnix(seconds int64) time.Time {
  return time.Unix(seconds, 0).In(sh)
}

func SHFrom(timeStr string, format string) (time.Time, error) {
  return time.ParseInLocation(format, timeStr, sh)
}

// 格式化string成int
func Atoi(s string) int {
  i, err := strconv.Atoi(s)
  if err != nil {
    return 0
  }
  return i
}

func ParseInDateTime(date string) time.Time {
  var dd, err = time.ParseInLocation(timeLayout24, date, sh)
  if err != nil {
    fmt.Printf("ParseInLocation [%s] error %v \n", date, err)
  }
  return dd
}

func ParseInDate(date string) time.Time {
  var dd, err = time.ParseInLocation(timeLayoutDate, date, sh)
  if err != nil {
    fmt.Printf("ParseInLocation [%s] error %v \n", date, err)
  }
  return dd
}

func FormatInDateTime(date time.Time) string {
  return date.Format(timeLayout24)
}

func FormatInDate(date time.Time) string {
  return date.Format(timeLayoutDate)
}

// []string 去重复
func DeDuplicationString(ids []string) []string {
  var result = make([]string, 0)
  var temp = map[string]bool{}
  for _, v := range ids {
    if temp[v] {
      continue
    }
    temp[v] = true
    result = append(result, v)
  }
  return result
}

// []int 去重复
func DeDuplicationInt(ids []int) []int {
  var result = make([]int, 0)
  var temp = map[int]bool{}
  for _, v := range ids {
    if temp[v] {
      continue
    }
    temp[v] = true
    result = append(result, v)
  }
  return result
}

func TimeFromMilliseconds(milliseconds int64) time.Time {
  var zero time.Time
  if milliseconds == 0 {
    return zero
  }
  return time.Unix(0, milliseconds*int64(time.Millisecond))
}

func FormatTime(mtime time.Time) string {
  return mtime.Format("2006-01-02 15:04:05")
}

func FormatDate(mtime time.Time) string {
  return mtime.Format("2006-01-02")
}

func Struct2Map(obj interface{}) map[string]interface{} {
  t := reflect.TypeOf(obj)
  v := reflect.ValueOf(obj)

  var data = make(map[string]interface{})
  for i := 0; i < t.NumField(); i++ {
    data[t.Field(i).Name] = v.Field(i).Interface()
  }
  return data
}

func SortAndConcat(param map[string]string) string {
  // 对 key 进行升序排序
  sortedKeys := make([]string, 0)
  for k, _ := range param {
    sortedKeys = append(sortedKeys, k)
  }
  sort.Strings(sortedKeys)
  // 对 key=value 的键值对用 & 连接起来，略过空值
  sbSignStr := NewStringBuilder()
  for i, k := range sortedKeys {
    if param[k] != "" {
      sbSignStr.Append(strings.ToLower(k))
      sbSignStr.Append("=")
      sbSignStr.Append(param[k])
      if i != (len(sortedKeys) - 1) {
        sbSignStr.Append("&")
      }
    }
  }
  return sbSignStr.ToString()
}
