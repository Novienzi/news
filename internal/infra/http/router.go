package http

import (
	"log"
	"net/http"
	"news/config"
	"news/internal/infra/http/news"

	"github.com/labstack/echo/v4"
)

func newEchoRouter(
	e *echo.Echo, cfg *config.Config,
	newsHandler *news.NewsHandlers,
	address string,
) {
	v1 := e.Group("/v1/api")
	health := v1.Group("/health")
	news := v1.Group("/news")
	news.POST("/", newsHandler.Create)
	news.GET("/:news_id", newsHandler.GetNewsByID)
	news.GET("/", newsHandler.GetNews)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	log.Fatal(e.Start(address))

}
