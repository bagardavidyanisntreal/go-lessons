package app

import (
	"context"

	"github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/model/order"
)

//go:generate mockgen -source=app.go -destination app_mocks.go -package app .

type orders interface {
	UpdateOrderStatus(ctx context.Context, id int64, status order.Status) error
}

func NewApplication(orders orders) *Application {
	return &Application{orders: orders}
}

type Application struct {
	orders orders
}
