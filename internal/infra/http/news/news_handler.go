package news

import (
	"net/http"
	"news/internal/infra/http/response"
	"news/internal/models"
	"news/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NewsHandlers struct {
	uc usecase.NewsUsecase
}

// NewNewsHandlers News handlers constructor
func NewNewsHandlers(newsUC usecase.NewsUsecase) *NewsHandlers {
	return &NewsHandlers{uc: newsUC}
}

func (h *NewsHandlers) Create(c echo.Context) error {
	ctx := c.Request().Context()
	req := NewsRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.RespondError(err))
	}

	createdNews, err := h.uc.Create(ctx, models.News{
		Title:  req.Title,
		Body:   req.Body,
		Author: req.Author,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.RespondError(err))
	}

	return c.JSON(http.StatusOK, response.RespondSuccess(createdNews))
}

func (h *NewsHandlers) GetNewsByID(c echo.Context) error {
	newsID, err := strconv.Atoi(c.Param("news_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.RespondError(err))
	}

	data, err := h.uc.GetNewsByID(c.Request().Context(), newsID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.RespondError(err))
	}

	return c.JSON(http.StatusOK, response.RespondSuccess(data))
}

func (h *NewsHandlers) GetNews(c echo.Context) error {
	ctx := c.Request().Context()
	req := GetNewsRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.RespondError(err))
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.RespondError(err))
	}

	data, p, err := h.uc.GetNews(ctx, req.Limit, req.Page, req.Search, req.Author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.RespondError(err))
	}
	return c.JSON(http.StatusOK, response.ResponseFetch(data, p))
}
