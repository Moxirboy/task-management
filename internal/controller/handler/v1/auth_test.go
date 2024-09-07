package v1

import (
	"github.com/gin-gonic/gin"
	"task-management/internal/service/usecase"
	"task-management/pkg/logger"
	"testing"
)

func TestAuthHandler_login(t *testing.T) {
	type fields struct {
		log       logger.Logger
		AccountUC usecase.IAccountUseCase
		AuthUC    usecase.IAuthUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &AuthHandler{
				log:       tt.fields.log,
				AccountUC: tt.fields.AccountUC,
				AuthUC:    tt.fields.AuthUC,
			}
			h.login(tt.args.c)
		})
	}
}

func TestAuthHandler_renew(t *testing.T) {
	type fields struct {
		log       logger.Logger
		AccountUC usecase.IAccountUseCase
		AuthUC    usecase.IAuthUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthHandler{
				log:       tt.fields.log,
				AccountUC: tt.fields.AccountUC,
				AuthUC:    tt.fields.AuthUC,
			}
			a.renew(tt.args.c)
		})
	}
}

func TestAuthHandler_signUp(t *testing.T) {
	type fields struct {
		log       logger.Logger
		AccountUC usecase.IAccountUseCase
		AuthUC    usecase.IAuthUseCase
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &AuthHandler{
				log:       tt.fields.log,
				AccountUC: tt.fields.AccountUC,
				AuthUC:    tt.fields.AuthUC,
			}
			h.signUp(tt.args.c)
		})
	}
}

func TestNewAuthHandler(t *testing.T) {
	type args struct {
		r         *gin.RouterGroup
		l         logger.Logger
		AuthUC    usecase.IAuthUseCase
		AccountUC usecase.IAccountUseCase
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewAuthHandler(tt.args.r, tt.args.l, tt.args.AuthUC, tt.args.AccountUC)
		})
	}
}
