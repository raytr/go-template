package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raytr/go-template/internal/model"
)

// ParsePaginationParams extracts pagination parameters from Gin context
func ParsePaginationParams(c *gin.Context) (*model.PaginationRequest, error) {
	// Check if required query parameters exist
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	if pageStr == "" {
		return nil, fmt.Errorf("page parameter is required")
	}

	if pageSizeStr == "" {
		return nil, fmt.Errorf("page_size parameter is required")
	}

	var pagination model.PaginationRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&pagination); err != nil {
		return nil, fmt.Errorf("invalid pagination parameters: %w", err)
	}

	// Additional validation
	if err := pagination.Validate(); err != nil {
		return nil, err
	}

	return &pagination, nil
}

// BuildPaginatedAPIResponse creates a standardized paginated API response
func BuildPaginatedAPIResponse(data interface{}, pagination *model.PaginationRequest, totalCount int64) gin.H {
	paginationResponse := model.NewPaginationResponse(pagination.Page, pagination.PageSize, totalCount)

	return gin.H{
		"data":       data,
		"pagination": paginationResponse,
	}
}
