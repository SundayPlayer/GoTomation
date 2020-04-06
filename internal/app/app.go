package app

import (
	"../config"
	"log"
	"os"
	"time"
)

var (
	logger *log.Logger
)

func Init(config *config.GlobalConfig) {
	f, err := os.OpenFile(config.App.Logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logger = log.New(f, "\033[1;34m" + time.Now().Format(time.RFC822) + " :\033[0m ", 0)

	for _, value := range config.Server {
		logger.Printf(
			"Server %s => User:%s, IP Address: %s, Port: %s\n", value.Name, value.User, value.IpAddr, value.Port)
	}
}
