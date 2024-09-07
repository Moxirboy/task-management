package middleware

import (
	"github.com/gin-gonic/gin"
)

type MyContext struct {
	*gin.Context

	id   string
	role string
}

func (c *MyContext) SetID(value string) {
	c.id = value
}

func (c *MyContext) SetRole(value string) {
	c.role = value
}

func (c *MyContext) GetID() string {
	return c.id
}

func (c *MyContext) GetRole() string {
	return c.role
}
