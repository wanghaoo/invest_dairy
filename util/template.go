package util

import (
  "html/template"
  "bytes"
  "io/ioutil"
)

//将参数导入html模版返回string
func ParsedTemplateToStr(templateName, fileName string, dataParam interface{}) (string, error) {
  t := template.New(templateName)
  t, err := t.ParseFiles(fileName)
  if err != nil {
    return "", err
  }
  buf := new(bytes.Buffer)
  err = t.Execute(buf, dataParam)
  if err != nil {
    return "", err
  }
  data,err := ioutil.ReadAll(buf)

  return string(data), nil
}
