package service

import (
	"github.com/Gircen/go-library-api/main/config"
	"log"
	"net/http"
)

var HttpServer = http.NewServeMux()

func run(config *config.Config) {
	HttpServer.HandleFunc("/healthCheck", HealthCheck)
	log.Printf(config.Server.Host + ":" + "3000")
	err := http.ListenAndServe(config.Server.Host+":"+"3000", HttpServer)
	if err != nil {
		println(err)
	}
}

func RunServiceCore(config *config.Config) {

	println("run")
	go run(config)
	println("service stopped")
}
