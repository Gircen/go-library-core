package service

import (
	"github.com/Gircen/go-library-api/main/config"
	"github.com/Gircen/go-library-core/main/logs"
	"net/http"
)

var HttpServer = http.NewServeMux()

func run(config *config.Config) {
	HttpServer.HandleFunc("/healthCheck", HealthCheck)
	logs.Logger.Printf(config.Server.Host + ":" + "3000")
	err := http.ListenAndServe(config.Server.Host+":"+"3000", HttpServer)
	if err != nil {
		println(err)
	}
}

func RunServiceCore(config *config.Config) {

	go logs.Log(config, &logs.Logger)
	logs.Logger.Println("run")
	go run(config)
	logs.Logger.Println("service stopped")
}
