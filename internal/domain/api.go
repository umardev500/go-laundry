package domain

// DTO

type PaginationInfo struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type DataWithPagination struct {
	Items      any            `json:"items"`
	Pagination PaginationInfo `json:"pagination"`
}

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"error,omitempty"`
}
