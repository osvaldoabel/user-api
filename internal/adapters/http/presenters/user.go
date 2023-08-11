package presenters

import (
	"time"

	"github.com/osvaldoabel/user-api/pkg/common"
)

type UserPresenter struct {
	ID      common.ID `json:"id"`
	Name    string    `json:"name"`
	Age     int       ` json:"age"`
	Email   string    ` json:"email"`
	Address string    `json:"address"`
	Status  string    ` json:"status"`

	CreatedAt time.Time ` json:"created_at"`
	UpdatedAt time.Time ` json:"updated_at"`
}

type UserCollection []UserPresenter

func NewUserCollection() UserCollection
