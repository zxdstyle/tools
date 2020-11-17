package role

import (
	"tools/app/models"
)

type PaginatorResponse struct {
	CurrentPage int            `json:"current_page"`
	Data        []models.Roles `json:"data"`
	PageSize    int            `json:"page_size"`
	Total       int64          `json:"total"`
}
