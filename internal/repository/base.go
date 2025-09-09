package repository

import (
	"fmt"

	"github.com/raytr/go-template/internal/model"
	"gorm.io/gorm"
)

// BasePaginationMethods provides common pagination functionality for GORM repositories
type BasePaginationMethods struct {
	db *gorm.DB
}

// NewPaginationMethods creates a new base pagination methods instance
func NewPaginationMethods(db *gorm.DB) *BasePaginationMethods {
	return &BasePaginationMethods{
		db: db,
	}
}

// ApplyPagination applies pagination parameters to a GORM query
func (b *BasePaginationMethods) ApplyPagination(query *gorm.DB, pagination *model.PaginationRequest) *gorm.DB {
	if pagination == nil {
		return query
	}

	offset := pagination.CalculateOffset()
	limit := pagination.PageSize

	return query.Limit(limit).Offset(offset)
}

// CountRecords counts total records for a given model
func (b *BasePaginationMethods) CountRecords(model interface{}) (int64, error) {
	var count int64

	if err := b.db.Model(model).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count records: %w", err)
	}

	return count, nil
}

// GetPaginatedRecords retrieves records with pagination
func (b *BasePaginationMethods) GetPaginatedRecords(
	dest interface{},
	model interface{},
	pagination *model.PaginationRequest,
	orderBy string,
) error {
	query := b.db.Model(model)

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	query = b.ApplyPagination(query, pagination)

	if err := query.Find(dest).Error; err != nil {
		return fmt.Errorf("failed to get paginated records: %w", err)
	}

	return nil
}
