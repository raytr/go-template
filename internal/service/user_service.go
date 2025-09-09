package service

import (
	"strings"

	"github.com/raytr/go-template/internal/model"
	"github.com/raytr/go-template/internal/repository"
)

// UserService handles business logic for users
type UserService struct {
	userRepo *repository.UserRepository
	*BasePaginationService
}

// NewUserService creates a new user service
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo:              userRepo,
		BasePaginationService: NewBasePaginationService(),
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(req *model.CreateUserReq) (*model.UserEntity, error) {
	// Create user entity
	user := &model.UserEntity{
		Code:    req.Code,
		Name:    req.Name,
		Email:   strings.ToLower(req.Email),
		Phone:   req.Phone,
		Address: req.Address,
	}

	// Save to database
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id uint) (*model.UserEntity, error) {
	return s.userRepo.GetByID(id)
}

// GetAllUsers retrieves all users with pagination
func (s *UserService) GetAllUsers(page, pageSize int) ([]*model.UserEntity, int64, error) {
	// Create and validate pagination request
	pagination, err := s.CreatePaginationRequest(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// Get users using new pagination system
	users, err := s.userRepo.GetAll(pagination)
	if err != nil {
		return nil, 0, err
	}

	// Get total count
	totalCount, err := s.userRepo.Count()
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id uint, req *model.UpdateUserReq) (*model.UserEntity, error) {
	// Check if user exists
	existingUser, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		existingUser.Name = req.Name
	}

	if req.Email != "" {
		existingUser.Email = strings.ToLower(req.Email)
	}

	if req.Phone != "" {
		existingUser.Phone = req.Phone
	}

	if req.Address != "" {
		existingUser.Address = req.Address
	}

	// Update user in database
	if err := s.userRepo.Update(existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
