package main

import (
	"io/ioutil"

	"../../internal/app"
	"../../internal/config"
)

func main() {
	fileContent, err := ioutil.ReadFile("gotomation.json")
	if err != nil {
		panic(err)
	}

	configs, err := config.Parse(fileContent)
	if err != nil {
		panic(err)
	}

	app.Init(configs)
}
