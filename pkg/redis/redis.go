package redis

import (
	"context"
	"encoding/json"
	"news/internal/models"
	"time"

	"github.com/go-redis/redis/v8"
)

type newsRedisRepo struct {
	redisClient *redis.Client
}

func New(c *redis.Client) *newsRedisRepo {
	client := &newsRedisRepo{
		redisClient: c,
	}
	return client
}

// Get new by id
func (n *newsRedisRepo) GetNewsByIDCtx(ctx context.Context, key string) (*models.News, error) {
	newsBytes, err := n.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	newsBase := &models.News{}
	if err = json.Unmarshal(newsBytes, newsBase); err != nil {
		return nil, err
	}

	return newsBase, nil
}

// Cache news item
func (n *newsRedisRepo) SetNewsCtx(ctx context.Context, key string, seconds int, news *models.News) error {
	newsBytes, err := json.Marshal(news)
	if err != nil {
		return err
	}
	if err = n.redisClient.Set(ctx, key, newsBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return err
	}
	return nil
}
