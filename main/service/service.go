package service

import (
	"github.com/Gircen/go-library-core/main/conf"
	"github.com/Gircen/go-library-core/main/logs"
	"net/http"
)

var HttpServer = http.NewServeMux()

func run() {
	HttpServer.HandleFunc("/healthCheck", HealthCheck)
	logs.Debug(conf.Conf.Server.Host + ":" + "3000")
	err := http.ListenAndServe(conf.Conf.Server.Host+":"+"3000", HttpServer)
	if err != nil {
		logs.Error(err)
	}
}

func Init() {
	conf.GetConfig()
	logs.Logger = logs.Log()
	go run()
}
