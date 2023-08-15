package entity

import (
	"time"

	"github.com/osvaldoabel/user-api/internal/dto"
	"github.com/osvaldoabel/user-api/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

type Status int

const (
	STATUS_INACTIVE Status = iota
	STATUS_ACTIVE
)

func (s Password) String() string {
	return string(s)
}

func (e Email) String() string {
	return string(e)
}

func (s Status) String() string {
	switch s {
	case STATUS_INACTIVE:
		return "INACTIVE"
	case STATUS_ACTIVE:
		return "ACTIVE"
	}
	return "UNKNOWN"
}

type User struct {
	ID       common.ID `valid:"uuid" gorm:"type:uuid;primary_key" json:"id"`
	Name     string    `json:"name"`
	Age      int       ` json:"age"`
	Email    Email     ` json:"email"`
	Password Password  ` json:"-"`
	Address  string    `json:"address"`
	Status   int       ` json:"status"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserEntity(dto dto.CreateUserInput) (*User, error) {
	var user User
	if dto.Password != "" {
		hash, err := NewPassword(dto.Password)
		if err != nil {
			return nil, err
		}

		user.Password = hash
	}

	user.ID = common.NewID()
	user.Email = Email(dto.Email)
	user.Name = dto.Name
	user.Age = dto.Age
	user.Address = dto.Address
	user.Status = int(STATUS_ACTIVE)

	return &user, nil

	// hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }

	// user := &User{
	// 	ID:       common.NewID(),
	// 	Name:     dto.Name,
	// 	Email:    Email(dto.Email),
	// 	Age:      dto.Age,
	// 	Address:  dto.Address,
	// 	Status:   int(STATUS_ACTIVE),
	// 	Password: Password(hash),
	// }

}

func (u *User) Update(dto dto.UpdateUserInput) error {
	if dto.Name != "" {
		u.Name = dto.Name
	}

	if dto.Age != 0 {
		u.Age = dto.Age
	}

	if dto.Address != "" {
		u.Address = dto.Address

	}

	if dto.Password != "" {
		pass, err := NewPassword(dto.Password)
		if err != nil {
			return err
		}

		u.Password = pass
	}

	return nil
}

func (u User) ValidatePassword(pass string) bool {
	_, err := NewPassword(pass)
	if err != nil {
		return false
	}

	// chekcs if it is the same passord after the encryuptation
	err = bcrypt.CompareHashAndPassword([]byte(u.Password.String()), []byte(pass))
	return (err == nil)
}
