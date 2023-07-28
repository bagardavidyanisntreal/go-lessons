package order

import "errors"

type Order struct {
	id     int64
	status Status
}

var ErrStatusDowngrade = errors.New("cannot downgrade status")

func (o *Order) SetStatus(status Status) error {
	if status < o.status {
		return ErrStatusDowngrade
	}

	o.status = status
	return nil
}

func (o *Order) ID() int64 {
	return o.id
}
