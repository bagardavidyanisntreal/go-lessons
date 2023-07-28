package imdb

import (
	"context"
	"sync"
	"time"

	"github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/model/cached"
	"github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/model/order"
)

//go:generate mockgen -source=order.go -destination order_mocks.go -package imdb .

type next interface {
	OrderByID(ctx context.Context, id int64) (*order.Order, error)
	UpdateOrder(ctx context.Context, order *order.Order) error
}

func NewOrder(next next) *Order {
	return &Order{
		next: next,
	}
}

type Order struct {
	lock sync.RWMutex
	ttl  time.Duration
	next next
	data map[int64]*cached.Cached[*order.Order]
}

func (o *Order) OrderByID(ctx context.Context, id int64) (*order.Order, error) {
	o.lock.RLock()
	cachedOrder, ok := o.data[id]
	o.lock.RUnlock()

	var (
		err       error
		orderByID *order.Order
	)

	now := time.Now()
	if !ok || cachedOrder.Expired(now) {
		orderByID, err = o.next.OrderByID(ctx, id)
		if err != nil {
			return nil, err
		}

		cachedOrder = cached.NewCached(orderByID, now.Add(o.ttl))

		o.lock.Lock()
		o.data[id] = cachedOrder
		o.lock.Unlock()
	}

	return cachedOrder.Data(), nil
}

func (o *Order) UpdateOrder(ctx context.Context, order *order.Order) error {
	err := o.next.UpdateOrder(ctx, order)
	if err != nil {
		return err
	}

	cachedOrder := cached.NewCached(order, time.Now().Add(o.ttl))

	o.lock.Lock()
	o.data[order.ID()] = cachedOrder
	o.lock.Unlock()

	return nil
}
