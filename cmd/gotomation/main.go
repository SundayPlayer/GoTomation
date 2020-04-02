package main

import (
	"../../internal/config"
	"io/ioutil"
)

func main() {
	fileContent, err := ioutil.ReadFile("gotomation.json")
	if err != nil {
		panic(err)
	}
	if err := config.Init(fileContent); err != nil {
		panic(err)
	}
}
