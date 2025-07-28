package logs

import (
	"github.com/Gircen/go-library-api/main/config"
	"log"
	"os"
)

var Logger log.Logger

func Log(config *config.Config, logger *log.Logger) {
	f, err := os.OpenFile("log/"+config.App.Name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			println(err)
		}
	}(f)

	logger = log.New(f, config.App.Name, log.LstdFlags)

}
