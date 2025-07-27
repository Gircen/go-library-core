package logs

import (
	"github.com/Gircen/go-library-api/main/config"
	"log"
	"os"
)

func Log(config *config.Config, logger *log.Logger) {
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Создаем логгер, который пишет в файл
	logger = log.New(f, config.App.Name, log.LstdFlags)

}
