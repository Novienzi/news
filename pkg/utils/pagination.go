package utils

import "news/internal/infra/http/response"

func BoundlessPaginationPageOffset(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	return limit, offset
}

func PaginationRes(page, count, limit int) response.PaginationVM {
	lastPage := count / limit
	if count%limit > 0 {
		lastPage = lastPage + 1
	}

	pagination := response.PaginationVM{
		CurrentPage:   page,
		LastPage:      lastPage,
		Count:         count,
		RecordPerPage: limit,
	}
	return pagination
}
