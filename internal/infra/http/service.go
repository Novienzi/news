package http

import (
	"news/config"
	"news/internal/infra/http/news"
	"news/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func NewProcess(
	cfg *config.Config,
	address string,
	newsUC usecase.NewsUsecase,
) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	newsHandler := news.NewNewsHandlers(newsUC)
	newEchoRouter(e, cfg,
		newsHandler, address,
	)
}
