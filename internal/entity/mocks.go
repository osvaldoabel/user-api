package entity

import "github.com/osvaldoabel/user-api/internal/dto"

func CreateSampleUser() dto.CreateUserInput {
	return dto.CreateUserInput{
		Name:     "user 01",
		Email:    "email@example.com",
		Password: "123456",
		Age:      20,
		Address:  "address 01",
	}
}
