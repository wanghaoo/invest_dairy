package common

import (
  "fmt"
  "github.com/Masterminds/squirrel"
  "strings"
)

func PrintQuery(query squirrel.SelectBuilder) {
  funcLog := Mlog.WithField("func", "PrintQuery")
  str, args, err := query.ToSql()
  if err != nil {
    Mlog.Error(err)
  }
  var msg []string
  msg = append(msg, str)
  if len(args) > 0 {
    msg = append(msg, "args:")
    for i, v := range args {
      msg = append(msg, fmt.Sprintf("%d => %v", i, v))
    }
  }
  msg = append(msg, "END  <<")
  funcLog.Debugf(strings.Join(msg, "\n"))
}

func PrintQueryWithSLog(query squirrel.SelectBuilder) {
  funcLog := Mlog.WithField("func", "PrintQueryWithSLog")
  str, args, err := query.ToSql()
  if err != nil {
    funcLog.Error(err)
  }
  var msg []string
  msg = append(msg, str)
  if len(args) > 0 {
    msg = append(msg, "args:")
    for i, v := range args {
      msg = append(msg, fmt.Sprintf("%d => %v", i, v))
    }
  }
  msg = append(msg, "END  <<")
  funcLog.Debugf(strings.Join(msg, "\n"))
}

func PrintUpdate(query squirrel.UpdateBuilder) {
  funcLog := Mlog.WithField("func", "PrintUpdate")
  str, args, err := query.ToSql()
  if err != nil {
    funcLog.Error(err)
  }
  var msg []string
  msg = append(msg, str)
  if len(args) > 0 {
    msg = append(msg, "args:")
    for i, v := range args {
      msg = append(msg, fmt.Sprintf("%d => %v", i, v))
    }
  }
  funcLog.Debugf(strings.Join(msg, "\n"))
}

func PrintInsert(query squirrel.InsertBuilder) {
  funcLog := Mlog.WithField("func", "PrintInsert")
  str, args, err := query.ToSql()
  if err != nil {
    funcLog.Error(err)
  }
  var msg []string
  //msg = append(msg, fmt.Sprintf("START >> SQL in [%s.%s]", log["file"], log["method"]))
  msg = append(msg, str)
  if len(args) > 0 {
    msg = append(msg, "args:")
    for i, v := range args {
      msg = append(msg, fmt.Sprintf("%d => %v", i, v))
    }
  }
  //msg = append(msg, "END  <<")
  funcLog.Debugf(strings.Join(msg, "\n"))
}

func PrintDelete(query squirrel.DeleteBuilder) {
  funcLog := Mlog.WithField("func", "PrintDelete")
  str, args, err := query.ToSql()
  if err != nil {
    funcLog.Error(err)
  }
  var msg []string
  //msg = append(msg, fmt.Sprintf("START >> SQL in [%s.%s]", log["file"], log["method"]))
  msg = append(msg, str)
  if len(args) > 0 {
    msg = append(msg, "args:")
    for i, v := range args {
      msg = append(msg, fmt.Sprintf("%d => %v", i, v))
    }
  }
  //msg = append(msg, "END  <<")
  funcLog.Debugf(strings.Join(msg, "\n"))
}
