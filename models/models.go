package models

import "gorm.io/gorm"

type IdModel struct {
	gorm.Model
}

type APIResponseBase struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Success bool        `json:"success"`
}

type APIResponse struct {
	APIResponseBase
}

type APIResponsePaginate struct {
	APIResponseBase
	Pagination
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	PageSize    int `json:"page_size"`
}
