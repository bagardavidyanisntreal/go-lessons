// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package app is a generated GoMock package.
package app

import (
	context "context"
	reflect "reflect"

	"github.com/bagardavidyanisntreal/go-lessons/goappdesignintro/example/model/order"
	gomock "github.com/golang/mock/gomock"
)

// Mockorders is a mock of orders interface.
type Mockorders struct {
	ctrl     *gomock.Controller
	recorder *MockordersMockRecorder
}

// MockordersMockRecorder is the mock recorder for Mockorders.
type MockordersMockRecorder struct {
	mock *Mockorders
}

// NewMockorders creates a new mock instance.
func NewMockorders(ctrl *gomock.Controller) *Mockorders {
	mock := &Mockorders{ctrl: ctrl}
	mock.recorder = &MockordersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockorders) EXPECT() *MockordersMockRecorder {
	return m.recorder
}

// UpdateOrderStatus mocks base method.
func (m *Mockorders) UpdateOrderStatus(ctx context.Context, id int64, status order.Status) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", ctx, id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockordersMockRecorder) UpdateOrderStatus(ctx, id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*Mockorders)(nil).UpdateOrderStatus), ctx, id, status)
}
