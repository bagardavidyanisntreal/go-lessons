package order

import (
	"context"
	"fmt"

	order2 "github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/model/order"
)

//go:generate mockgen -source=order.go -destination order_mocks.go -package order .

type orderer interface {
	OrderByID(ctx context.Context, id int64) (*order2.Order, error)
	UpdateOrder(ctx context.Context, order *order2.Order) error
}

func NewService(orderer orderer) *Service {
	return &Service{
		orderer: orderer,
	}
}

type Service struct {
	orderer orderer
}

const logFormat = "[order service]"

func (s *Service) UpdateOrderStatus(ctx context.Context, id int64, status order2.Status) error {
	orderByID, err := s.orderer.OrderByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%s order request error %w", logFormat, err)
	}

	// notify other dependencies about status changing

	err = orderByID.SetStatus(status)
	if err != nil {
		return fmt.Errorf("%s status updating fail %w", logFormat, err)
	}

	err = s.orderer.UpdateOrder(ctx, orderByID)
	if err != nil {
		return fmt.Errorf("%s order update error %w", logFormat, err)
	}

	return nil
}
