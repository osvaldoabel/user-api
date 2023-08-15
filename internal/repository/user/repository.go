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
func (ur *userDBRepository) FindAll(params entity.Pagination) (*entity.Pagination, error) {
	var users []entity.User

	// Get total rows
	var totalRows int64
	ur.Db.Model(entity.User{}).Count(&totalRows)

	// Get total pages
	totalPages := totalRows / int64(params.Limit)

	result := ur.Db.Limit(params.Limit).
		Offset(params.Page).
		Order(DEFAULT_ORDER_BY).
		Find(&users)

	if result.Error != nil {
		// log here
		return nil, result.Error
	}

	_pagination := entity.Pagination{
		Limit:      params.Limit,
		Page:       params.Page,
		TotalRows:  totalRows,
		TotalPages: int(totalPages),
		Rows:       users,
	}
	return &_pagination, nil
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

	return user, nil
}

// Delete
func (ur *userDBRepository) Delete(id entity.ID) error {
	var user entity.User
	return ur.Db.Delete(&user, "id = ?", id).Error
}
