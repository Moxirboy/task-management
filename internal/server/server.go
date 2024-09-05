package server

import (
	"fmt"
	"task-management/internal/configs"
	"task-management/pkg/logger"
)

type Server struct {
	cfg    *configs.Config
	logger logger.Logger
}

func NewServer(
	cfg *configs.Config,
	logger logger.Logger,
) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s Server) Run() error {

	g := gin.New()
	g.Use(gin.Recovery())
	return g.Run(fmt.Sprintf(":%d", s.cfg.Server.Port))

}
