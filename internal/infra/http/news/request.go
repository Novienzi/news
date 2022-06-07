package news

type NewsRequest struct {
	Author string `json:"author,omitempty" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" db:"content" validate:"required"`
}

type GetNewsRequest struct {
	Page   int    `query:"page" validate:"required"`
	Limit  int    `query:"limit" validate:"required"`
	Author string `query:"author"`
	Search string `query:"q"`
}
