package http_server

import "time"

type Option func(*Server)

func WithPort(port string) Option {
	return func(server *Server) {
		server.server.Addr = port
	}
}

func WithReadTimeout(t time.Duration) Option {
	return func(server *Server) {
		server.server.ReadTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) Option {
	return func(server *Server) {
		server.server.WriteTimeout = t
	}
}

func WithShutdownTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.shutdownTimeout = timeout
	}
}
