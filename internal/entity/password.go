package entity

import "errors"

type Password string

func NewPassword(pass string) (Password, error) {
	// checks if it's not in the blacklist passwords
	if IsUnauthorizedPass(pass) {
		return Password(""), errors.New("unauthorized password")
	}

	return Password(pass), nil
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
	return []Password{"123", "abc"}
}

func (u User) UpdatePassword(pass string) (Password, error) {
	newPass, err := NewPassword(pass)
	if err != nil {
		return Password(pass), err
	}

	return Password(newPass), nil
}
