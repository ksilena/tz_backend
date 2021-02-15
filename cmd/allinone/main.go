package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"tz/httpserver"
	"tz/server"
)

type ServerConfig struct {
	Name  string
	Limit int
}

type Config struct {
	ClientDir     string
	ListenAddress string
	Servers       []ServerConfig
}

func main() {

	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	err = json.NewDecoder(bytes.NewReader(b)).Decode(&config)
	if err != nil {
		panic(err)
	}

	hs := httpserver.New(config.ClientDir, config.ListenAddress)
	var nextServerInt server.Input = hs
	for _, c := range config.Servers {

		srv := server.New(c.Name, c.Limit)
		srv.SetInput(nextServerInt)
		srv.Start()
		nextServerInt = srv
	}

	hs.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
