package stdrouter

import "github.com/VanOODev/education_web_server/adapters/controllers/http/stdrouter/config"

type Server struct {
}

func NewServer(cfg config.Config) *Server {
	return &Server{}
}

func (s *Server) Close() {
}
