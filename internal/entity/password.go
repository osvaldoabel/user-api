package entity

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	PASSWORD_BLACKLIST = []Password{"123", "abc"}
)

type Password string

func NewPassword(pass string) (Password, error) {
	// checks if it's not in the blacklist passwords
	if len(pass) < 3 {
		return Password(""), errors.New("password must be at least 3 characters")
	}

	// checks if it's not in the blacklist passwords
	if IsUnauthorizedPass(pass) {
		return Password(""), errors.New("unauthorized password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return Password(""), err
	}

	return Password(hash), nil
}

func IsUnauthorizedPass(pass string) bool {
	for _, unauthorizedPass := range GetPasswordBlackList() {
		if pass == unauthorizedPass.String() {
			return true
		}
	}

	return false
}

func GetPasswordBlackList() []Password {
	return PASSWORD_BLACKLIST
}

func UpdatePassword(pass string) (Password, error) {
	newPass, err := NewPassword(pass)
	if err != nil {
		return Password(pass), err
	}

	return newPass, nil
}
