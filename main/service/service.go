package service

import (
	"github.com/Gircen/go-library-api/main/config"
	"github.com/Gircen/go-library-core/main/logs"
	"log"
	"net/http"
)

func run(config *config.Config) {
	http.HandleFunc("/healthCheck", HealthCheck)
	log.Printf(config.Server.Host + ":" + config.Server.Port)
	err := http.ListenAndServe(config.Server.Host+":"+config.Server.Port, nil)
	if err != nil {
		println(err)
	}
}

func RunServiceCore(config *config.Config) {
	go logs.ReadLog()
	println("run")
	run(config)
	println("service started")
}
