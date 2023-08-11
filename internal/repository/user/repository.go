package user

import (
	"github.com/osvaldoabel/user-api/internal/entity"
	"gorm.io/gorm"
)

const (
	DEFAULT_ORDER_BY = "created_at desc"
)

type userDBRepository struct {
	Db *gorm.DB
}

func NewUserDBRepository(db *gorm.DB) UserDBRepository {
	return &userDBRepository{
		Db: db,
	}
}

// FindAll
func (ur *userDBRepository) FindAll(params entity.Pagination) ([]entity.User, error) {
	var users []entity.User

	result := ur.Db.Limit(params.Limit).
		Offset(params.Offset).
		Order(DEFAULT_ORDER_BY).
		Find(&users)

	if result.Error != nil {
		// log here
		return []entity.User{}, result.Error
	}

	return users, nil
}

// FindByID
func (ur *userDBRepository) FindByID(id entity.ID) (entity.User, error) {
	var user entity.User
	result := ur.Db.Find(&user, "id=?", id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (ur *userDBRepository) FindByEmail(email entity.Email) (entity.User, error) {
	var user entity.User
	result := ur.Db.Find(&user, "email=?", email.String())
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

// Insert
func (ur *userDBRepository) Insert(user entity.User) (entity.User, error) {
	result := ur.Db.Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

// Update
func (ur *userDBRepository) Update(user entity.User) (entity.User, error) {
	err := ur.Db.Save(&user).Error
	if err != nil {
		// log here
		return entity.User{}, err
	}
	return entity.User{}, nil
}

// Delete
func (ur *userDBRepository) Delete(id entity.ID) error {
	var user entity.User
	return ur.Db.Delete(&user, "id = ?", id).Error
}
