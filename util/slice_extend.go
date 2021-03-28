package util

import (
	"strconv"
	"strings"
)

func StringSliceToInt(arr []string) ([]int, error) {
	result := make([]int, 0)
	for _, v := range arr {
	  if len(v) <= 0 {
	    continue
    }
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		} else {
			result = append(result, i)
		}
	}
	return result, nil
}

func IntSliceToString(arr []int) []string {
  result := make([]string, 0)
  for _, v := range arr {
    result = append(result, strconv.Itoa(v))
  }
  return result
}

func CommaSeparatedIntSlice(str string) ([]int, error) {
	return StringSliceToInt(strings.Split(str, ","))
}

func IntSliceToCommaSeparated(ids []int) string {
	var array = make([]string, 0)
	for _, id := range ids {
		array = append(array, strconv.Itoa(id))
	}
	return strings.Join(array, ",")
}

func KeysOfMapIntBool(m map[int]bool) []int {
	keys := make([]int, 0, len(m))
	for k, v := range m {
		if v {
			keys = append(keys, k)
		}
	}
	return keys
}

func RemoveDupliatedInt(arr []int) []int {
	var m = make(map[int]bool, 0)
	for _, v := range arr {
		m[v] = true
	}
	return KeysOfMapIntBool(m)
}

func KeysOfMapStringBool(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k, v := range m {
		if v {
			keys = append(keys, k)
		}
	}
	return keys
}

func RemoveDupliatedString(arr []string) []string {
	var m = make(map[string]bool, 0)
	for _, v := range arr {
		m[v] = true
	}
	return KeysOfMapStringBool(m)
}

func SplitArticleImages(outputImages string) []string {
  if CheckStrIsBlank(outputImages) {
    return []string{}
  }
  return strings.Split(outputImages, ",")
}

// 检测重复值 并返回该值
func VerifySliceDuplicates(slice []interface{}) interface{} {
  tempMap := map[interface{}]int{}
  for _, v := range slice {
    if _, ok := tempMap[v]; ok {
      return v
    }
    tempMap[v] = 0
  }
  return nil
}
