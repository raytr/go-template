package service

import (
	"github.com/raytr/go-template/internal/model"
)

// BasePaginationService provides common pagination functionality for all services
type BasePaginationService struct{}

// NewBasePaginationService creates a new base pagination service
func NewBasePaginationService() *BasePaginationService {
	return &BasePaginationService{}
}

// CreatePaginationRequest creates a pagination request from raw parameters
func (b *BasePaginationService) CreatePaginationRequest(page, pageSize int) (*model.PaginationRequest, error) {
	pagination := &model.PaginationRequest{
		Page:     page,
		PageSize: pageSize,
	}

	if err := pagination.Validate(); err != nil {
		return nil, err
	}

	return pagination, nil
}
