package repository

import (
	"errors"
	"fmt"

	"github.com/raytr/go-template/internal/model"
	"gorm.io/gorm"
)

// UserRepository handles database operations for users using GORM
type UserRepository struct {
	db *gorm.DB
	*BasePaginationMethods
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:                    db,
		BasePaginationMethods: NewPaginationMethods(db),
	}
}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *model.UserEntity) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id uint) (*model.UserEntity, error) {
	var user model.UserEntity

	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// GetAll retrieves all users with pagination
func (r *UserRepository) GetAll(pagination *model.PaginationRequest) ([]*model.UserEntity, error) {
	var users []*model.UserEntity

	if err := r.GetPaginatedRecords(&users, &model.UserEntity{}, pagination, "created_at DESC"); err != nil {
		return nil, err
	}

	return users, nil
}

// Update updates an existing user
func (r *UserRepository) Update(user *model.UserEntity) error {
	if err := r.db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete removes a user from the database
func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&model.UserEntity{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// Count returns the total number of users
func (r *UserRepository) Count() (int64, error) {
	return r.CountRecords(&model.UserEntity{})
}
