package usecase

import (
	"context"
	"news/internal/models"
	"news/internal/models/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mdlMock := &mocks.NewsRepository{}
	rdsMock := &mocks.RedisRepository{}

	uc := newsUc{
		mdl:     mdlMock,
		rdsRepo: rdsMock,
	}

	var (
		ctx           = context.Background()
		newsID        = 1
		newsAuthor    = "vien"
		newsTitle     = "Gunung Api Meletus"
		newsBody      = "Pada hari senin, Gunung Api Purba di daerah z meletus dahsyat"
		newsCreatedAt = time.Now()

		news = models.News{
			ID:        newsID,
			Author:    newsAuthor,
			Title:     newsTitle,
			Body:      newsBody,
			CreatedAt: newsCreatedAt,
		}
	)

	t.Run("success case", func(t *testing.T) {
		mdlMock.On("Create", ctx, models.News{
			Author: newsAuthor,
			Title:  newsTitle,
			Body:   newsBody,
		}).Return(news, nil).Once()

		res, err := uc.Create(ctx, models.News{
			Author: newsAuthor,
			Title:  newsTitle,
			Body:   newsBody,
		})
		assert.NoError(t, err)
		assert.Equal(t, news, res)
	})
}
