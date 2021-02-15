package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"tz/httpclient"
	"tz/httpserver"
	"tz/server"
)

type Config struct {
	ServerName    string
	ClientDir     string
	Limit         int
	NextServer    *string
	ListenAddress string
}

func main() {
	var filename = flag.String("c", "serverconfig1.json", "config file path")
	flag.Parse()
	b, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}
	var config Config
	err = json.NewDecoder(bytes.NewReader(b)).Decode(&config)
	if err != nil {
		panic(err)
	}
	s := server.New(config.ServerName, config.Limit)

	i := httpserver.New(config.ClientDir, config.ListenAddress)
	s.SetInput(i)
	if config.NextServer != nil {
		hc := httpclient.New(*config.NextServer)
		hc.SetInput(s)
		hc.Start()
	}

	i.Start()
	s.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
