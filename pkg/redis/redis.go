package redis

import (
	"context"
	"fmt"
	"food-delivery/internal/configs"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var (
	instance *redis.Client
	once     sync.Once
)

// DB return database connection
func DB(cfg *configs.Redis) (*redis.Client, error) {
	var err error
	once.Do(func() {
		fmt.Println(cfg)
		instance = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Username: "default",
			Password: cfg.Password,
			DB:       cfg.DB, // default DB
		})

	})
	err = instance.Ping(context.Background()).Err()
	if err != nil {
		return nil, errors.Wrap(err, "redis.Connect")
	}

	return instance, nil
}
