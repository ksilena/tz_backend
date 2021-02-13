package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// Config структура конфигурационного файла http-сервера
type Config struct {
	Name          string
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

}
