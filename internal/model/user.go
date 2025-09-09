package model

import (
	"time"
)

// UserEntity represents the users table in the database
type UserEntity struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email"`
	Phone     string    `gorm:"type:varchar(50)" json:"phone,omitempty"`
	Address   string    `gorm:"type:text" json:"address,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName specifies the table name for UserEntity
func (UserEntity) TableName() string {
	return "users"
}

// CreateUserReq represents the request for creating a new user
type CreateUserReq struct {
	Code    string `json:"code" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
}

// UpdateUserReq represents the request for updating a user
type UpdateUserReq struct {
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty" binding:"omitempty,email"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
}

// UserResponse represents the response for user data
type UserResponse struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	Address   string    `json:"address,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToResponse converts UserEntity to UserResponse
func (u *UserEntity) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Code:      u.Code,
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Address:   u.Address,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
