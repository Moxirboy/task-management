package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strings"
	cnst "task-management/internal/const"
	"task-management/internal/dto"
	"task-management/internal/models"
	"task-management/pkg/errors"
)

// CasbinMiddleware is the middleware function for checking permissions using Casbin.
func (m Middleware) CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := &MyContext{Context: c}
		pass, err := m.CheckPermission(mc)
		if err == nil && pass {
			c.Next()
			return
		}

		if err == nil {
			c.JSON(http.StatusForbidden, dto.BaseResponse{
				Status: string(cnst.ErrorStatusForbidden),
			})
			c.Abort()
			return
		}

		c.JSON(errors.Parse(err))
		c.Abort()
	}
}

// CheckPermission checks if the user has the necessary permissions.
func (m Middleware) CheckPermission(c *MyContext) (bool, error) {
	user, err := m.getUserName(c)
	if err != nil {
		return false, err
	}

	method := c.Request.Method
	path := c.Request.URL.Path

	roles := strings.Split(user, ",")
	for _, role := range roles {
		ok, err := m.enforcer.Enforce(path, role, method)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}

	return false, nil
}

// getUserName retrieves the user's role based on the access token.
func (m Middleware) getUserName(c *MyContext) (string, error) {
	accessToken := c.GetHeader("Authorization")
	jwt := c.GetHeader("JWT")
	fmt.Println(jwt)
	fmt.Println(accessToken)
	if accessToken == "" {
		return string(models.UNAUTHORIZED), nil
	}

	if check := strings.Split(accessToken, " "); len(check) == 2 &&
		strings.EqualFold(check[0], string(models.BASIC)) {
		return string(models.BASIC), nil
	}

	token, err := m.getToken(accessToken)
	if err != nil {
		return "", err
	}

	if token == nil {
		token, err = m.checkAuth(accessToken)
		if err != nil {
			return "", err
		}
		err = m.redisRepo.SetAccessToken(
			accessToken,
			token,
		)
		if err != nil {
			return "", err
		}
	}

	c.SetID(token.ID)
	c.SetRole(token.Role)
	c.Set("ID", token.ID)

	return token.Role, nil
}

// getToken retrieves the access token from Redis.
func (m Middleware) getToken(accessToken string) (*models.AccessToken, error) {
	token, err := m.redisRepo.AccessToken(accessToken)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return token, nil
}

// checkAuth validates the access token and retrieves the user ID and role.
func (m Middleware) checkAuth(accessToken string) (*models.AccessToken, error) {
	id, role, err := m.authUseCase.Check(
		context.Background(),
		accessToken,
	)
	if err != nil {
		return nil, err
	}

	return &models.AccessToken{
		ID:   id,
		Role: role,
	}, nil
}
