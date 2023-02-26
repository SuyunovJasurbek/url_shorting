package storage

import (
	"context"
	"time"

	"github.com/SuyunovJasurbek/url_shorting/models"
)

type StorageI interface {
	User() UserI
	Url() UrlI
}

type UserI interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, urerID string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}

type UrlI interface {
	CreateUrl(ctx context.Context, url *models.Url) (*models.Url, error)
	DeleteUrl(ctx context.Context, urlID string) (*models.Url, error)
	GetUrlByID(ctx context.Context, urlID string) (*models.Url, error)
	GetUrls(ctx context.Context, url string) (*models.GetAllUrl, error)
	UpdateUrl(ctx context.Context, url *models.Url) (*models.Url, error)
}

type CacheStorageI interface {
	Redis() RedisI
}
type RedisI interface {
	Set(ctx context.Context, key, value string, expTime time.Duration) error
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (value string, err error)
	Contains(ctx context.Context, key string) (bool, error)
}
