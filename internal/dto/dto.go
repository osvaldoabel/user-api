package dto

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Status   string `json:"status"`
}

type UpdateUserInput struct {
	ID       string `validate:"notnull" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Status   string `json:"status"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}

type ListUsers struct {
}
