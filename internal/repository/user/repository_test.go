package user

import (
	"testing"

	"github.com/osvaldoabel/user-api/internal/dto"
	"github.com/osvaldoabel/user-api/internal/entity"
	"github.com/osvaldoabel/user-api/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	db := database.NewDbTest()
	userRepo := NewUserDBRepository(db.DB)

	userDTO := entity.CreateSampleUser()
	user, err := entity.NewUserEntity(userDTO)
	assert.Nil(t, err)
	assert.NotNil(t, *user)

	userResult, err := userRepo.Insert(*user)
	assert.Nil(t, err)
	assert.NotNil(t, userResult)
}

func TestCreateUser_Error(t *testing.T) {
	db := database.NewInmemoryDB()
	userRepo := NewUserDBRepository(db)

	userDTO := entity.CreateSampleUser()
	user, _ := entity.NewUserEntity(userDTO)

	_, err := userRepo.Insert(*user)
	assert.NotNil(t, err)
}

func TestUpdateUser_Success(t *testing.T) {
	db := database.NewDbTest()
	userRepo := NewUserDBRepository(db.DB)

	userDTO := entity.CreateSampleUser()
	user, err := entity.NewUserEntity(userDTO)
	assert.Nil(t, err)
	assert.NotNil(t, *user)

	userResult, err := userRepo.Insert(*user)
	assert.Nil(t, err)
	assert.NotNil(t, userResult)

	userResult.Update(dto.UpdateUserInput{
		Name:    "new User 02",
		Age:     90,
		Address: "new Address 02 ",
	})
	updated, err := userRepo.Update(userResult)
	assert.Nil(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, "new User 02", updated.Name)
}

func TestUpdateUser_Error(t *testing.T) {
	db := database.NewDbTest()
	userRepo := NewUserDBRepository(db.DB)

	userDTO := entity.CreateSampleUser()
	user, err := entity.NewUserEntity(userDTO)
	assert.Nil(t, err)
	assert.NotNil(t, *user)

	userResult, err := userRepo.Insert(*user)
	assert.Nil(t, err)
	assert.NotNil(t, userResult)

	userResult.Update(dto.UpdateUserInput{
		Name:    "new User 02",
		Age:     90,
		Address: "new Address 02 ",
	})

	// update connection to force error
	simpleConnection := database.NewInmemoryDB()
	userRepo = NewUserDBRepository(simpleConnection)

	_, err = userRepo.Update(userResult)
	assert.NotNil(t, err)

}

func TestDeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		db := database.NewDbTest()
		userRepo := NewUserDBRepository(db.DB)

		userDTO := entity.CreateSampleUser()
		user, err := entity.NewUserEntity(userDTO)
		assert.Nil(t, err)
		assert.NotNil(t, *user)

		userResult, err := userRepo.Insert(*user)
		assert.Nil(t, err)
		assert.NotNil(t, userResult)

		err = userRepo.Delete(entity.ID(userResult.ID.String()))
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		db := database.NewInmemoryDB()
		userRepo := NewUserDBRepository(db)
		err := userRepo.Delete(entity.ID("1234567890-not-exists"))
		assert.NotNil(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		db := database.NewDbTest()
		userRepo := NewUserDBRepository(db.DB)

		userDTO := entity.CreateSampleUser()
		user, err := entity.NewUserEntity(userDTO)
		assert.Nil(t, err)
		assert.NotNil(t, *user)

		userResult, err := userRepo.Insert(*user)
		assert.Nil(t, err)
		assert.NotNil(t, userResult)

		//user 02
		userDTO2 := entity.CreateSampleUser()
		userDTO2.Name = "user 02"

		user2, err := entity.NewUserEntity(userDTO)
		userResult2, err := userRepo.Insert(*user2)
		assert.Nil(t, err)
		assert.NotNil(t, userResult2)

		//test find all
		users, err := userRepo.FindAll(entity.Pagination{
			Limit: 10,
			Page:  0,
		})

		assert.Nil(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 2, len(users.Rows.([]entity.User)))
	})
}
