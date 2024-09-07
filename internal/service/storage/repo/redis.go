package repo

import "task-management/internal/models"

type IRedisRepository interface {
	AccessToken(key string) (*models.AccessToken, error)
	SetAccessToken(key string, accessToken *models.AccessToken) error
}
