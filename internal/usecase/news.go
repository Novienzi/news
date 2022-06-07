package usecase

import (
	"context"
	"news/internal/infra/http/response"
	"news/internal/models"
	"news/pkg/utils"
	"strconv"
)

// News use case
type NewsUsecase interface {
	Create(ctx context.Context, news models.News) (models.News, error)
	GetNewsByID(ctx context.Context, newsID int) (*models.News, error)
	GetNews(ctx context.Context, limit, page int, q, author string) (*[]models.News, response.PaginationVM, error)
}

type newsUc struct {
	mdl     models.NewsRepository
	rdsRepo models.RedisRepository
}

var cacheDuration = 3600

func NewNewsUsecase(m models.NewsRepository, r models.RedisRepository) NewsUsecase {
	return &newsUc{
		mdl:     m,
		rdsRepo: r,
	}
}

func (uc *newsUc) Create(ctx context.Context, param models.News) (models.News, error) {
	res, err := uc.mdl.Create(ctx, param)
	if err != nil {
		return res, err
	}
	return res, err
}
func (uc *newsUc) GetNewsByID(ctx context.Context, newsID int) (*models.News, error) {
	newsBase, err := uc.rdsRepo.GetNewsByIDCtx(ctx, strconv.Itoa(newsID))
	if err != nil {
		return nil, err
	}
	if newsBase != nil {
		return newsBase, nil
	}

	data, err := uc.mdl.GetNewsByID(ctx, newsID)
	if err != nil {
		return nil, err
	}

	if err = uc.rdsRepo.SetNewsCtx(ctx, strconv.Itoa(newsID), cacheDuration, data); err != nil {
		return nil, err
	}

	return data, err
}
func (uc *newsUc) GetNews(ctx context.Context, limit, page int, q, author string) (*[]models.News, response.PaginationVM, error) {
	size, offset := utils.BoundlessPaginationPageOffset(page, limit)
	p := response.PaginationVM{}

	res, err := uc.mdl.GetNews(ctx, size, offset, q, author)
	if err != nil {
		return nil, p, err
	}
	count := len(res)
	if limit > 0 {
		p = utils.PaginationRes(page, count, limit)
	}
	return &res, p, nil
}
