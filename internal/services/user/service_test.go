package user

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/osvaldoabel/user-api/internal/container"
	"github.com/osvaldoabel/user-api/internal/entity"
	repomock "github.com/osvaldoabel/user-api/internal/repository/user/mock"
	"github.com/osvaldoabel/user-api/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := repomock.NewMockUserDBRepository(ctrl)
		userRepo.EXPECT().Insert(gomock.Any()).Return(entity.User{}, nil)

		dps := container.DependencyContainer{
			UserRepo: userRepo,
		}

		userService := NewUserService(dps)
		result, err := userService.CreateUser(entity.User{}, context.Background())
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		userRepo := repomock.NewMockUserDBRepository(ctrl)
		userRepo.EXPECT().Insert(gomock.Any()).Return(entity.User{}, errors.New("database error"))

		dps := container.DependencyContainer{
			UserRepo: userRepo,
		}

		userService := NewUserService(dps)
		_, err := userService.CreateUser(entity.User{}, context.Background())
		assert.NotNil(t, err)
	})
}

func TestListUsers(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uuid := common.ID(uuid.New())

		userRepo := repomock.NewMockUserDBRepository(ctrl)
		userRepo.EXPECT().FindAll(entity.Pagination{
			Limit: 1,
			Page:  1,
		}).Return(&entity.Pagination{
			Limit: 1,
			Page:  1,
			Rows: []entity.User{
				{
					ID:    uuid,
					Name:  "user 01",
					Email: "",
				},
			},
		}, nil)

		dps := container.DependencyContainer{
			UserRepo: userRepo,
		}

		userService := NewUserService(dps)
		result, err := userService.ListUsers(entity.Pagination{
			Limit: 1,
			Page:  1,
		}, context.Background())
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, len(result.Rows.([]entity.User)))
	})
}
