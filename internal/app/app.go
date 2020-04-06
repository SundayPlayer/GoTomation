package app

import (
	"../config"
	"log"
	"os"
	"time"
)

type App struct {
	Config config.GlobalConfig
	Log    *log.Logger
}

func Init(config config.GlobalConfig) {
	app := &App{
		Config: config,
	}
	f, err := os.OpenFile(app.Config.App.Logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	app.Log = log.New(f, "\033[1;34m" + time.Now().Format(time.RFC822) + " :\033[0m ", 0)

	for _, value := range app.Config.Server {
		app.Log.Printf(
			"Server %s => User:%s, IP Address: %s, Port: %s\n", value.Name, value.User, value.IpAddr, value.Port)
	}
}
