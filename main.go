package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/ksilena/tz_backend.git/httpserver"
)

// Config структура конфигурационного файла http-сервера
type Config struct {
	Name          string
	Path          string
	ListenAddress string
	Limit         int
}

func main() {

	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	err = json.NewDecoder(bytes.NewReader(file)).Decode(&config)
	if err != nil {
		panic(err)
	}

	httpserver := httpserver.New(config.Path, config.ListenAddress)
	httpserver.Start()
}
