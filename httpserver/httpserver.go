package httpserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpServer struct {
	out           chan []byte
	uploadPath    string
	listenAddress string
}

func New(path string, listenAddress string) *HttpServer {
	s := &HttpServer{}
	s.uploadPath = path
	s.listenAddress = listenAddress
	return s
}

func (s *HttpServer) Start() {
	fs := http.FileServer(http.Dir(s.uploadPath))
	http.Handle("/client/", http.StripPrefix("", fs))
	http.HandleFunc("/upload", s.upload)
	fmt.Println(http.Dir(s.uploadPath))

	go http.ListenAndServe(s.listenAddress, nil)
}
func (i *HttpServer) upload(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	i.out <- b

}

func (s *HttpServer) SetOutputChannel(ch chan []byte) {
	s.out = ch
}
