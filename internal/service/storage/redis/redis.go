package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	accessTokenExpireTime = time.Minute * 2
)

type redisRepository struct {
	rds *redis.Client
}

func NewRedisRepository(rds *redis.Client) repo.IRedisRepository {
	return &redisRepository{rds: rds}
}

func (r redisRepository) AccessToken(
	key string,
) (*models.AccessToken, error) {
	var token models.AccessToken
	if err := r.rds.Get(
		context.Background(),
		key,
	).Scan(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (r redisRepository) SetAccessToken(
	key string,
	accessToken *models.AccessToken,
) error {
	err := r.rds.Set(
		context.Background(),
		key,
		accessToken,
		accessTokenExpireTime,
	).Err()

	return err
}
