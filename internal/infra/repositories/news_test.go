package repositories_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"news/internal/infra/repositories"
	"news/internal/models"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
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

	getAllNews = []models.News{
		{
			ID:        newsID,
			Author:    newsAuthor,
			Title:     newsTitle,
			Body:      newsBody,
			CreatedAt: newsCreatedAt,
		},
	}
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreate(t *testing.T) {

	db, mock := NewMock()
	repo := repositories.NewNewsRepository(db)
	defer func() {
		db.Close()
	}()

	query := `
	INSERT INTO news (author, title, body) 
	VALUES ($1, $2, $3) 
	RETURNING *
	`
	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "author", "title", "body", "created_at"}).AddRow(newsID, newsAuthor, newsTitle, newsBody, newsCreatedAt)

		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(newsAuthor, newsTitle, newsBody).
			WillReturnRows(rows)

		res, err := repo.Create(context.Background(), news)
		assert.NoError(t, err)
		assert.Equal(t, news, res)
	})

	t.Run("error exec", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WillReturnError(fmt.Errorf("exec error"))
		mock.ExpectRollback()

		_, err := repo.Create(context.Background(), news)
		assert.Error(t, err)
	})

}

func TestGetNewsByID(t *testing.T) {
	db, mock := NewMock()
	repo := repositories.NewNewsRepository(db)
	defer func() {
		db.Close()
	}()

	query := `
	SELECT * FROM news where id = $1
	`

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "author", "title", "body", "created_at"}).AddRow(newsID, newsAuthor, newsTitle, newsBody, newsCreatedAt)

		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(newsID).
			WillReturnRows(rows)

		res, err := repo.GetNewsByID(context.Background(), newsID)
		assert.NoError(t, err)
		assert.Equal(t, &news, res)
	})

	t.Run("error exec", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WillReturnError(fmt.Errorf("exec error"))
		mock.ExpectRollback()

		_, err := repo.GetNewsByID(context.Background(), newsID)
		assert.Error(t, err)
	})

}

func TestGetNews(t *testing.T) {
	db, mock := NewMock()
	repo := repositories.NewNewsRepository(db)
	defer func() {
		db.Close()
	}()

	query := `
	SELECT * FROM news n WHERE (n.title ILIKE $1 OR n.body ILIKE $2) AND n.author ILIKE $3 LIMIT $4 OFFSET $5
	`

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "author", "title", "body", "created_at"}).AddRow(newsID, newsAuthor, newsTitle, newsBody, newsCreatedAt)

		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("%keyword%", "%keyword%", "%author%", 10, 1).
			WillReturnRows(rows)

		res, err := repo.GetNews(context.Background(), 10, 1, "keyword", "author")
		assert.NoError(t, err)
		assert.Equal(t, getAllNews, res)
	})

	t.Run("error exec", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WillReturnError(fmt.Errorf("exec error"))
		mock.ExpectRollback()

		_, err := repo.GetNews(context.Background(), 10, 0, "keyword", "author")
		assert.Error(t, err)
	})

}
