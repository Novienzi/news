// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "news/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// NewsRepository is an autogenerated mock type for the NewsRepository type
type NewsRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, news
func (_m *NewsRepository) Create(ctx context.Context, news models.News) (models.News, error) {
	ret := _m.Called(ctx, news)

	var r0 models.News
	if rf, ok := ret.Get(0).(func(context.Context, models.News) models.News); ok {
		r0 = rf(ctx, news)
	} else {
		r0 = ret.Get(0).(models.News)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.News) error); ok {
		r1 = rf(ctx, news)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNews provides a mock function with given fields: ctx, limit, offset, q, author
func (_m *NewsRepository) GetNews(ctx context.Context, limit int, offset int, q string, author string) ([]models.News, error) {
	ret := _m.Called(ctx, limit, offset, q, author)

	var r0 []models.News
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string, string) []models.News); ok {
		r0 = rf(ctx, limit, offset, q, author)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, string, string) error); ok {
		r1 = rf(ctx, limit, offset, q, author)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNewsByID provides a mock function with given fields: ctx, newsID
func (_m *NewsRepository) GetNewsByID(ctx context.Context, newsID int) (*models.News, error) {
	ret := _m.Called(ctx, newsID)

	var r0 *models.News
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.News); ok {
		r0 = rf(ctx, newsID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, newsID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
