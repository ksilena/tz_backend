package httpclient

import (
	"bytes"
	"fmt"
	"net/http"
)

type Input interface {
	SetOutputChannel(chan []byte)
}

type HttpClient struct {
	in         chan []byte
	nextServer string
}

func New(nextServer string) *HttpClient {
	return &HttpClient{
		nextServer: nextServer,
	}
}

func (s *HttpClient) Start() {
	go func() {

		for b := range s.in {
			resp, err := http.Post(s.nextServer, "application/json", bytes.NewBuffer(b))
			if err != nil {
				fmt.Println(err)
				continue
			}
			resp.Body.Close()
		}

	}()

}

func (s *HttpClient) SetInput(i Input) {
	ch := make(chan []byte, 1)
	s.in = ch
	i.SetOutputChannel(ch)
	return
}
