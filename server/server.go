package server

import (
	"github.com/JanusDG/real-time-group-chat/server/session"
)

type Server struct {
	session session.Session
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) HandleConnection() bool {
	return false
}

func (s *Server) CreateSession() bool {
	return false
}

func (s *Server) RunSession() bool {
	return false
}
