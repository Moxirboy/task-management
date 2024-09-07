package server

import (
	"context"
	"fmt"
	"food-delivery/internal/configs"
	"food-delivery/internal/controller/handler"
	"food-delivery/internal/controller/middleware"
	"food-delivery/internal/models"
	rds "food-delivery/internal/service/storage/redis"
	"food-delivery/internal/service/usecase"
	"food-delivery/pkg/logger"
	"food-delivery/pkg/postgres"
	"food-delivery/pkg/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	rDB, err := redis.DB(&s.cfg.Redis)
	if err != nil {
		s.logger.Fatal(err)
	}
	pDB, err := postgres.DB(&s.cfg.Postgres)
	if err != nil {
		s.logger.Fatal(err)
	}

	g := gin.New()
	g.Use(gin.Recovery())
	uc := usecase.New(s.cfg, pDB, s.logger)

	redisRe := rds.NewRedisRepository(rDB)
	middleware.SetUpMiddleware(
		g,
		s.cfg,
		rds.NewRedisRepository(rDB),
		uc.IAuthUseCase(),
	)
	s.logger.Info(s.cfg)
	handler.SetUp(&g.RouterGroup, s.cfg, uc, s.logger, redisRe)

	uc.IAccountUseCase().CreateUser(context.Background(), &models.User{
		FirstName: s.cfg.Setup.AdminName,
		LastName:  s.cfg.Setup.AdminLastName,
		Email:     s.cfg.Setup.AdminEmail,
		Position:  string(models.PositionAdmin),
		Password:  s.cfg.Setup.AdminPassword,
	})
	return g.Run(fmt.Sprintf(":%d", s.cfg.Server.Port))

}
