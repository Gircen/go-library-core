package logs

import (
	logEnum "github.com/Gircen/go-library-api/main/log"
	"github.com/Gircen/go-library-core/main/conf"
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func Warn(v ...any) {
	if logEnum.IsLog(conf.Conf.Log.Level, logEnum.WARN) {
		Logger.Println(logEnum.WARN.String(), v)
	}
}

func Error(v ...any) {
	Logger.Println(logEnum.ERROR.String(), v)
}

func Info(v ...any) {
	Logger.Println(logEnum.INFO.String(), v)
}

func Debug(v ...any) {
	Logger.Println(logEnum.DEBUG.String(), v)
}

func Log() *log.Logger {

	err := os.MkdirAll("log/", os.ModePerm)
	if err != nil {
		return nil
	}
	f, err := os.OpenFile("log/"+conf.Conf.App.Name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	var logger = log.New(f, conf.Conf.App.Name, log.LstdFlags|log.Llongfile)
	mw := io.MultiWriter(os.Stdout, f)
	logger.SetOutput(mw)
	return logger
}
