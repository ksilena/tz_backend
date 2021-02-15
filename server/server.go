package server

import (
	"encoding/json"
	"fmt"
	"time"
	"tz/limiter"
	"tz/packet"
)

type Input interface {
	SetOutputChannel(chan []byte)
}

type Starter interface {
	New(name string, limit int) *Server
	Start()
	SetInput(i Input)
}

type Server struct {
	in    chan []byte
	Name  string
	out   chan []byte
	limit int
}

func New(name string, limit int) *Server {
	return &Server{
		Name:  name,
		limit: limit,
	}
}

func (s *Server) Start() {
	lim := limiter.New(s.limit)
	go func() {
		var p []packet.Packet
		for b := range s.in {
			if !lim.Check() {
				if s.out != nil {
					s.out <- b
					continue
				}

				fmt.Printf("%s: Превышен лимит кластера!\n", s.Name)

			}

			json.Unmarshal(b, &p)

			s.log(&p)
		}

	}()

}

func (s *Server) SetInput(i Input) {
	ch := make(chan []byte, 1)
	s.in = ch
	i.SetOutputChannel(ch)
	return
}

func (s *Server) SetOutputChannel(ch chan []byte) {
	s.out = ch
}

func (s *Server) log(p *[]packet.Packet) {
	time := time.Now().Format("01-02-2006 15:04:05")
	fmt.Println(time, " Сервер: ", s.Name, "packet ", *p)
}
