package repositories

import (
	"context"
	"database/sql"
	"news/internal/models"
	"news/pkg/utils"
	"strings"
)

type newsRepo struct {
	db *sql.DB
}

func NewNewsRepository(db *sql.DB) models.NewsRepository {
	return &newsRepo{db: db}
}

func (model *newsRepo) Create(ctx context.Context, param models.News) (res models.News, err error) {
	query := `
	INSERT INTO news (author, title, body) 
	VALUES ($1, $2, $3) 
	RETURNING *
	`
	err = model.db.QueryRowContext(ctx, query, param.Author, param.Title, param.Body).Scan(&res.ID,
		&res.Author, &res.Title, &res.Body, &res.CreatedAt)
	if err != nil {
		return
	}

	return res, err
}

func (model *newsRepo) GetNewsByID(ctx context.Context, newsID int) (*models.News, error) {
	res := models.News{}
	query := `
	SELECT * FROM news where id = $1`
	err := model.db.QueryRowContext(ctx, query, newsID).Scan(&res.ID, &res.Author, &res.Title, &res.Body, &res.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &res, err
}
func (model *newsRepo) GetNews(ctx context.Context, limit, offset int, q, author string) (res []models.News, err error) {
	selectQuery := `SELECT * FROM news n`
	queryParams := []interface{}{}

	var whereQuery string
	if q != "" {
		keyword := "%" + q + "%"
		whereQuery += ` WHERE (n.title ILIKE ? OR n.body ILIKE ?)`
		queryParams = append(queryParams, keyword, keyword)
	}
	if author != "" {
		keyword := "%" + author + "%"
		whereQuery += ` WHERE n.author ILIKE ?`
		queryParams = append(queryParams, keyword)
	}

	if author != "" && q != "" {
		whereQuery = `WHERE (n.title ILIKE ? OR n.body ILIKE ?) AND n.author ILIKE ?`
	}

	limitQuery := `LIMIT ? OFFSET ?`
	queryParams = append(queryParams, limit, offset)

	finalQuery := strings.Join([]string{selectQuery, whereQuery, limitQuery}, " ")
	query := utils.SubstitutePlaceholder(finalQuery, 1)

	rows, err := model.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		resTemp := models.News{}
		err = rows.Scan(&resTemp.ID, &resTemp.Author, &resTemp.Title, &resTemp.Body, &resTemp.CreatedAt)
		if err != nil {
			return
		}
		res = append(res, resTemp)
	}
	return
}
