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
	if logEnum.WARN.IsLog(conf.Conf.Log.Level) {
		Logger.Println(logEnum.WARN.String(), v)
	}
}

func FatalError(v ...any) {
	if logEnum.ERROR.IsLog(conf.Conf.Log.Level) {
		Logger.Fatalln(logEnum.ERROR.String(), v)
	}
}

func Error(v ...any) {
	if logEnum.ERROR.IsLog(conf.Conf.Log.Level) {
		Logger.Println(logEnum.ERROR.String(), v)
	}
}

func Info(v ...any) {
	if logEnum.INFO.IsLog(conf.Conf.Log.Level) {
		Logger.Println(logEnum.INFO.String(), v)
	}
}

func Debug(v ...any) {
	if logEnum.DEBUG.IsLog(conf.Conf.Log.Level) {
		Logger.Println(logEnum.DEBUG.String(), v)
	}
}

func Log() {

	err := os.MkdirAll("log/", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("log/"+conf.Conf.App.Name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(f, conf.Conf.App.Name, log.LstdFlags|log.Llongfile)
	mw := io.MultiWriter(os.Stdout, f)
	Logger.SetOutput(mw)
}
