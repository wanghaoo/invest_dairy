package util

import (
  "bytes"
  "crypto/md5"
  "encoding/hex"
  "fmt"
  "regexp"
  "strconv"
  "strings"
)

func StringNotNull(input *string) string {
	if input != nil {
		return *input
	}
	return ""
}

//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
  rs := []rune(str)
  rl := len(rs)
  end := 0

  if start < 0 {
    start = rl - 1 + start
  }
  end = start + length

  if start > end {
    start, end = end, start
  }

  if start < 0 {
    start = 0
  }
  if start > rl {
    start = rl
  }
  if end < 0 {
    end = 0
  }
  if end > rl {
    end = rl
  }

  return string(rs[start:end])
}

// 获取中括号内容
func TrimBrackets(content string) string {
  reg := regexp.MustCompile("^【([^】]+)】")
  res := reg.FindStringSubmatch(content)
  if len(res) > 1 {
    return res[1]
  }
  return ""
}

// 获取中括号内容
func TrimBracketContent(content string) string {
  reg := regexp.MustCompile("^【([^】]+)】")
  return reg.ReplaceAllString(content, "")
}

// 去除空格／换行符
func TrimSpaceAndBreak(str string) string {
  // str = strings.Replace(str, " ", "", -1)
  // str = strings.Replace(str, "\n", "", -1)
  str = strings.Replace(str, "\u200b", "", -1)
  str = strings.Replace(str, "<em>", "*", -1)
  str = strings.Replace(str, "</em>", "*", -1)
  str = strings.TrimSpace(str)
  return str
}

func ConvertIntArrayToStringArray (intArray []int) ([]string) {
  result := make([]string, 0)
  for _, intValue := range intArray {
    result = append(result, strconv.Itoa(intValue))
  }
  return result
}

func Md5Str(str string) string {
  m := md5.New()
  m.Write([]byte(str))
  return hex.EncodeToString(m.Sum(nil))
}

type StringBuilder struct {
  buf bytes.Buffer
}

func (this *StringBuilder) Append(obj interface{}) *StringBuilder {
  this.buf.WriteString(fmt.Sprintf("%v", obj))
  return this
}

func NewStringBuilder() *StringBuilder {
  return &StringBuilder{buf: bytes.Buffer{}}
}

func (this *StringBuilder) ToString() string {
  return this.buf.String()
}

func GetName(name string)(string, string, string) {
  if len(name) <= 0 {
    return "", "", ""
  }
  names := strings.Split(name, " ")
  if len(names) > 2 {
    return names[0], names[1], names[2]
  }
  if len(names) == 2 {
    return names[0], "", names[1]
  }
  return names[0], "", ""
}