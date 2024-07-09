package httpserver

import (
	"net/http"
)

type ServerConfig struct {
	Port string
}

type Server struct {
	server *http.Server
	notify chan error
}

func NewServer(handler http.Handler, cfg *ServerConfig) *Server {
	httpServer := &http.Server{
		Handler: handler,
		Addr:    cfg.Port,
	}
	return &Server{
		server: httpServer,
		notify: make(chan error, 1),
	}
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}
