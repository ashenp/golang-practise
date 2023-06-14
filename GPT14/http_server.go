package GPT14

import (
	"net/http"
	"strconv"
)

type HttpServer struct {
	port int
	mux  *http.ServeMux
}

func (p *HttpServer) start() {
	http.ListenAndServe(":"+strconv.Itoa(p.port), p.mux)
}

func (p *HttpServer) RegisterFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	p.mux.HandleFunc(pattern, handler)
}

func NewHttpServer(port int) *HttpServer {
	return &HttpServer{
		port: port,
		mux:  http.NewServeMux(),
	}
}
