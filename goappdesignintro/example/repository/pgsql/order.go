package pgsql

import (
	"context"

	"github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/model/order"
)

func NewRepository() *Repository {
	return &Repository{}
}

type Repository struct {
}

func (o *Repository) OrderByID(ctx context.Context, id int64) (*order.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (o *Repository) UpdateOrder(ctx context.Context, order *order.Order) error {
	// TODO implement me
	panic("implement me")
}
