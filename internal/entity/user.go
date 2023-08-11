package entity

import (
	"time"

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

func NewUser(name string, email string, pass string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}

	return &User{
		ID:     common.NewID(),
		Name:   name,
		Email:  Email(email),
		Status: int(STATUS_ACTIVE),
		// Password: string(hash),
		Password: Password(hash),
	}, nil

}

func (u User) ValidatePassword(pass string) bool {
	passObj, err := NewPassword(pass)
	if err != nil {
		return false
	}

	// chekcs if it is the same passord after the encryuptation
	err = bcrypt.CompareHashAndPassword([]byte(u.Password.String()), []byte(passObj.String()))
	return (err == nil)
}
