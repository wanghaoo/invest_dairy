package common

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Mlog *logrus.Logger

func init() {
  Mlog = logrus.New()
  Mlog.SetReportCaller(true)
  Mlog.Formatter = &logrus.TextFormatter{
    CallerPrettyfier: func(f *runtime.Frame) (string, string) {
      filename := path.Base(f.File)
      return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
    },
  }
  Mlog.Level = logrus.DebugLevel
  file, err := os.OpenFile("autocue.log", os.O_CREATE|os.O_WRONLY, 0666)
  if err == nil {
    Mlog.Out = file
  } else {
    Mlog.Info("Failed to log to file, using default stderr")
  }
  mw := io.MultiWriter(os.Stdout, file)
  Mlog.SetOutput(mw)
  defer func() {
    err := recover()
    if err != nil {
      entry := err.(*logrus.Entry)
      Mlog.WithFields(logrus.Fields{
        "omg":         true,
        "err_animal":  entry.Data["animal"],
        "err_size":    entry.Data["size"],
        "err_level":   entry.Level,
        "err_message": entry.Message,
        "number":      100,
      }).Error("The ice breaks!") // or use Fatal() to force the process to exit with a nonzero code
    }
  }()
}