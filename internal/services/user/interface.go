package user

import (
	"context"

	"github.com/osvaldoabel/user-api/internal/dto"
	"github.com/osvaldoabel/user-api/internal/entity"
)

type UserService interface {
	// CreateUser
	CreateUser(user entity.User, ctx context.Context) (entity.User, error)

	// UpdateUser
	UpdateUser(dto dto.UpdateUserInput, ctx context.Context) (entity.User, error)

	// DeleteUser
	DeleteUser(userID entity.ID, ctx context.Context) error

	// FindUserByEmail
	FindUserByEmail(email entity.Email, ctx context.Context) (entity.User, error)

	// GetUser
	FindUser(userID entity.ID, ctx context.Context) (entity.User, error)

	// ListUsers
	ListUsers(search entity.Pagination, ctx context.Context) ([]entity.User, error)
}
