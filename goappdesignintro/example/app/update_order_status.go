package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	order2 "github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/model/order"
)

type UpdateOrderStatusRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

var errStatusUndefined = errors.New("status undefined")

func (r UpdateOrderStatusRequest) InternalStatus() (order2.Status, error) {
	switch r.Status {
	case "new":
		return order2.StatusNew, nil
	case "in-progress":
		return order2.StatusInProgress, nil
	case "fail":
		return order2.StatusFail, nil
	case "ready":
		return order2.StatusReady, nil
	}

	return -1, errStatusUndefined
}

func (a *Application) UpdateOrderStatus(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var req UpdateOrderStatusRequest
	err := decoder.Decode(&req)
	if err != nil {
		response500(writer, err)
		return
	}

	ctx := request.Context()
	status, err := req.InternalStatus()
	if err != nil {
		response400(writer, err)
		return
	}

	err = a.orders.UpdateOrderStatus(ctx, req.ID, status)
	if err != nil {
		if err == order2.ErrStatusDowngrade {
			response400(writer, err)
			return
		}

		response500(writer, err)
		return
	}

	response200(writer, fmt.Sprintf("order %d status updated %s", req.ID, req.Status))
}
