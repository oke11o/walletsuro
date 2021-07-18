package handler

type Server struct {
	service service
}

func NewServer(s service) *Server {
	return &Server{service: s}
}
