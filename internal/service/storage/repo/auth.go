package repo

import (
	"context"
	"time"
)

type IAuthTokenRepository interface {
	Create(ctx context.Context, token, id, role string, date time.Time) error
	CleanUp(ctx context.Context) error
	Get(ctx context.Context, token string) (*dto.AuthToken, error)
	Delete(ctx context.Context, token string) error
}
