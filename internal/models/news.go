package models

import (
	"context"
	"time"
)

// News base model
type News struct {
	ID        int       `json:"id" db:"id" validate:"required"`
	Author    string    `json:"author,omitempty" db:"author" validate:"required"`
	Title     string    `json:"title" db:"title" validate:"required,gte=10"`
	Body      string    `json:"body" db:"content" validate:"required,gte=20"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

type NewsRepository interface {
	Create(ctx context.Context, news News) (News, error)
	GetNewsByID(ctx context.Context, newsID int) (*News, error)
	GetNews(ctx context.Context, limit, offset int, q, author string) ([]News, error)
}

type RedisRepository interface {
	GetNewsByIDCtx(ctx context.Context, key string) (*News, error)
	SetNewsCtx(ctx context.Context, key string, seconds int, news *News) error
}
