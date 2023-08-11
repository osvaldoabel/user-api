package entity

import (
	"time"

	"github.com/osvaldoabel/user-api/pkg/common"
)

type Product struct {
	ID common.ID `valid:"uuid" gorm:"type:uuid;primary_key" json:"id"`

	Name  string `json:"name"`
	Price int    ` json:"price"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        common.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Vaildate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Vaildate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := common.ParseID(p.ID.String()); err != nil {
		return ErrIDIsInvalid
	}

	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if p.Price == 0 {
		return ErrInvalidPrice

	}

	return nil
}
