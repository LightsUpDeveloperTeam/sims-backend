package utils

import _ "time"

type Response struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      *Error      `json:"error,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Error struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Details *string `json:"details,omitempty"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	PerPage     int `json:"per_page"`
	TotalItems  int `json:"total_items"`
}
