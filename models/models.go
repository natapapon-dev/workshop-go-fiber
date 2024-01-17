package models

import "gorm.io/gorm"

type IdModel struct {
	gorm.Model
}

type APIResponseBase struct {
	Data    interface{}
	Message string
	Status  int
	Success bool
}

type APIResponse struct {
	APIResponseBase
}

type APIResponsePaginate struct {
	APIResponseBase
	CurrentPage int
	TotalPage   int
	PageSize    int
}
