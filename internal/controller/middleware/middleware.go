package middleware

import (
	"fmt"
	"log"
	"task-management/internal/configs"
	"task-management/internal/service/storage/repo"
	"task-management/internal/service/usecase"
	"time"

	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	app         *gin.Engine
	config      *configs.Config
	redisRepo   repo.IRedisRepository
	enforcer    *casbin.Enforcer
	authUseCase usecase.IAuthUseCase
}

func SetUpMiddleware(
	app *gin.Engine,
	config *configs.Config,
	redisRepo repo.IRedisRepository,
	authUseCase usecase.IAuthUseCase,
) {
	m := Middleware{
		app:         app,
		config:      config,
		redisRepo:   redisRepo,
		authUseCase: authUseCase,
	}

	m.ginMiddleware()

	enforcer, err := casbin.NewEnforcer(m.config.Casbin.ConfigPath, m.config.Casbin.Name)
	if err != nil {
		panic(err)
	}

	if err = enforcer.LoadPolicy(); err != nil {
		log.Println("LoadPolicy failed, err: ", err)
	}
	m.enforcer = enforcer
	roleManager, ok := m.enforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl)
	if !ok {
		log.Fatalf("Expected *defaultrolemanager.RoleManagerImpl, but got %T", m.enforcer.GetRoleManager())
	}

	roleManager.AddMatchingFunc("keyMatch", util.KeyMatch)
	roleManager.AddMatchingFunc("keyMatch2", util.KeyMatch2)

	m.enforcer.EnableLog(true)
	m.app.Use(m.CasbinMiddleware())
}

func (m Middleware) ginMiddleware() {
	m.app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s | status=%d | %s | %s | %s | in=%d | out=%d | uri=%s | error=%s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Request.ContentLength,
			param.BodySize,
			param.Path,
			param.ErrorMessage,
		)
	}))

	m.app.Use(gin.Recovery())
	m.app.Use(m.CORSMiddleware())
}

func (m Middleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
