package repo

import (
	"context"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, login, password string) (*models.User, error)
	CheckField(
		ctx context.Context,
		field, value string,
	) (bool, error)
}
