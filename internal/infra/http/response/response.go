package response

type ApiResponseModel struct {
	Success    bool         `json:"success"`
	Data       interface{}  `json:"data"`
	Error      string       `json:"error, omitempty"`
	Pagination PaginationVM `json:"pagination, omitempty"`
}

type PaginationVM struct {
	CurrentPage   int `json:"current_page"`
	LastPage      int `json:"last_page"`
	Count         int `json:"count"`
	RecordPerPage int `json:"record_per_page"`
}

func RespondSuccess(data interface{}) ApiResponseModel {
	return ApiResponseModel{
		Success: true,
		Data:    data,
	}
}

func RespondError(err error) ApiResponseModel {
	return ApiResponseModel{
		Success: false,
		Data:    nil,
		Error:   err.Error(),
	}
}

func ResponseFetch(data interface{}, pagination PaginationVM) ApiResponseModel {
	return ApiResponseModel{
		Success:    true,
		Data:       data,
		Pagination: pagination,
	}
}
