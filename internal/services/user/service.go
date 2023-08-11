package user

import (
	"context"

	"github.com/osvaldoabel/user-api/internal/container"
	"github.com/osvaldoabel/user-api/internal/dto"
	"github.com/osvaldoabel/user-api/internal/entity"
	"github.com/osvaldoabel/user-api/internal/repository/user"
)

type userService struct {
	UserRepo user.UserDBRepository
}

func NewUserService(deps container.DependencyContainer) UserService {
	return &userService{
		UserRepo: deps.UserRepo,
	}
}

// CreateUser
func (us *userService) CreateUser(user entity.User, ctx context.Context) (entity.User, error) {
	user, err := us.UserRepo.Insert(user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// UpdateUser
func (us *userService) UpdateUser(dto dto.UpdateUserInput, ctx context.Context) (entity.User, error) {
	user, err := us.UserRepo.FindByID(entity.ID(dto.ID))
	if err != nil {
		return entity.User{}, err
	}

	if dto.Name != "" {
		user.Name = dto.Name
	}

	if dto.Password != "" {
		dtoPass, err := entity.NewPassword(dto.Password)
		if err != nil {
			return entity.User{}, err
		}

		user.Password = dtoPass
	}

	updatedUser, err := us.UserRepo.Update(user)
	if err != nil {
		return entity.User{}, err
	}

	return updatedUser, nil
}

// DeleteUser
func (us *userService) DeleteUser(userID entity.ID, ctx context.Context) error {
	err := us.UserRepo.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) FindUser(userID entity.ID, ctx context.Context) (entity.User, error) {
	user, err := us.UserRepo.FindByID(userID)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// FindUserByEmail
func (us *userService) FindUserByEmail(userEmail entity.Email, ctx context.Context) (entity.User, error) {
	user, err := us.UserRepo.FindByEmail(userEmail)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// ListUsers
func (us *userService) ListUsers(search entity.Pagination, ctx context.Context) ([]entity.User, error) {
	users, err := us.UserRepo.FindAll(search)
	if err != nil {
		return []entity.User{}, err
	}

	return users, nil
}
