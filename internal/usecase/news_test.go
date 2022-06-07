package usecase

import (
	"context"
	"database/sql"
	"errors"
	"news/internal/models"
	"news/internal/models/mocks"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func TestCreate(t *testing.T) {
	mdlMock := &mocks.NewsRepository{}
	rdsMock := &mocks.RedisRepository{}

	uc := newsUc{
		mdl:     mdlMock,
		rdsRepo: rdsMock,
	}

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

	t.Run("error case", func(t *testing.T) {
		mdlMock.On("Create", ctx, models.News{
			Author: newsAuthor,
			Title:  newsTitle,
			Body:   newsBody,
		}).Return(models.News{}, sql.ErrNoRows).Once()

		_, err := uc.Create(ctx, models.News{
			Author: newsAuthor,
			Title:  newsTitle,
			Body:   newsBody,
		})
		assert.Error(t, err)
	})
}

func TestGetNewsByID(t *testing.T) {
	mdlMock := &mocks.NewsRepository{}
	rdsMock := &mocks.RedisRepository{}

	uc := newsUc{
		mdl:     mdlMock,
		rdsRepo: rdsMock,
	}

	t.Run("success case from db", func(t *testing.T) {
		rdsMock.On("GetNewsByIDCtx", ctx, strconv.Itoa(newsID)).Return(nil, nil).Once()
		mdlMock.On("GetNewsByID", ctx, newsID).Return(&news, nil)
		rdsMock.On("SetNewsCtx", ctx, strconv.Itoa(newsID), 3600, &news).Return(nil).Once()

		res, err := uc.GetNewsByID(ctx, newsID)
		assert.NoError(t, err)
		assert.Equal(t, &news, res)
	})

	t.Run("success case from redis", func(t *testing.T) {
		rdsMock.On("GetNewsByIDCtx", ctx, strconv.Itoa(newsID)).Return(&news, nil).Once()

		res, err := uc.GetNewsByID(ctx, newsID)
		assert.NoError(t, err)
		assert.Equal(t, &news, res)
	})

	t.Run("error case", func(t *testing.T) {
		rdsMock.On("GetNewsByIDCtx", ctx, strconv.Itoa(newsID)).Return(nil, nil).Once()
		mdlMock.On("GetNewsByID", ctx, newsID).Return(nil, sql.ErrNoRows).Once()
		rdsMock.On("SetNewsCtx", ctx, strconv.Itoa(newsID), 3600, &news).Return(errors.New("error")).Once()

		_, err := uc.GetNewsByID(ctx, newsID)
		assert.Error(t, err)
	})
}

func TestGetNews(t *testing.T) {
	mdlMock := &mocks.NewsRepository{}
	rdsMock := &mocks.RedisRepository{}

	uc := newsUc{
		mdl:     mdlMock,
		rdsRepo: rdsMock,
	}

	t.Run("success", func(t *testing.T) {
		mdlMock.On("GetNews", ctx, 10, 0, "keyword", "keyword").Return([]models.News{news}, nil)

		res, _, err := uc.GetNews(ctx, 10, 0, "keyword", "keyword")
		assert.NoError(t, err)
		assert.Equal(t, &[]models.News{news}, res)
	})
}
