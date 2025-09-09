package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raytr/go-template/internal/model"
	"github.com/raytr/go-template/internal/utils"
)

// PaginationHandler provides common pagination functionality for all handlers
type PaginationHandler struct{}

// NewPaginationHandler creates a new base pagination handler
func NewPaginationHandler() *PaginationHandler {
	return &PaginationHandler{}
}

// ParsePagination extracts and validates pagination parameters from Gin context
func (p *PaginationHandler) ParsePagination(c *gin.Context) (*model.PaginationRequest, error) {
	return utils.ParsePaginationParams(c)
}

// RespondWithPaginatedData sends a standardized paginated response
func (p *PaginationHandler) RespondWithPaginatedData(
	c *gin.Context,
	statusCode int,
	data interface{},
	pagination *model.PaginationRequest,
	totalCount int64,
) {
	response := utils.BuildPaginatedAPIResponse(data, pagination, totalCount)
	c.JSON(statusCode, response)
}

// RespondWithPaginationError sends a standardized error response for pagination errors
func (p *PaginationHandler) RespondWithPaginationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
