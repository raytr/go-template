package model

import (
	"fmt"
	"math"
)

// PaginationRequest represents pagination parameters for requests
type PaginationRequest struct {
	Page     int `json:"page" form:"page" binding:"required,min=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"required,min=1,max=100"`
}

// Validate validates pagination parameters
func (p *PaginationRequest) Validate() error {
	if p.Page < 1 {
		return fmt.Errorf("page is required and must be >= 1")
	}
	if p.PageSize < 1 || p.PageSize > 100 {
		return fmt.Errorf("page_size is required and must be between 1 and 100")
	}
	return nil
}

// CalculateOffset calculates the database offset
func (p *PaginationRequest) CalculateOffset() int {
	return (p.Page - 1) * p.PageSize
}

// PaginationResponse represents pagination metadata in responses
type PaginationResponse struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// NewPaginationResponse creates a new pagination response
func NewPaginationResponse(page, pageSize int, total int64) *PaginationResponse {
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if total == 0 {
		totalPages = 0
	}

	return &PaginationResponse{
		Page:       page,
		PageSize:   pageSize,
		Total:      int(total),
		TotalPages: totalPages,
	}
}
