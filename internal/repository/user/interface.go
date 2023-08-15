package user

import (
	"github.com/osvaldoabel/user-api/internal/entity"
)

type ReaderRepository interface {
	// GetAll resturns a list of Users
	// FindAll(params entity.Pagination) ([]entity.User, error)
	FindAll(params entity.Pagination) (*entity.Pagination, error)

	FindByEmail(email entity.Email) (entity.User, error)

	// FindByID one specific user filtering by ID
	FindByID(id entity.ID) (entity.User, error)
}

type WriterRepository interface {
	// Insert creates a new User record in the DB
	Insert(user entity.User) (entity.User, error)

	// Update updates an existing User already existent in the DB
	Update(user entity.User) (entity.User, error)

	// Delete
	Delete(id entity.ID) error
}

type UserDBRepository interface {
	WriterRepository
	ReaderRepository
}
