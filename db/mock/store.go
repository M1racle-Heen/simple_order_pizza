// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/M1racle-Heen/simple_order_pizza/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/M1racle-Heen/simple_order_pizza/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateCustomer mocks base method.
func (m *MockStore) CreateCustomer(arg0 context.Context, arg1 db.CreateCustomerParams) (db.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", arg0, arg1)
	ret0, _ := ret[0].(db.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCustomer indicates an expected call of CreateCustomer.
func (mr *MockStoreMockRecorder) CreateCustomer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockStore)(nil).CreateCustomer), arg0, arg1)
}

// CreateOrder mocks base method.
func (m *MockStore) CreateOrder(arg0 context.Context, arg1 db.CreateOrderParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockStoreMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockStore)(nil).CreateOrder), arg0, arg1)
}

// CreatePayment mocks base method.
func (m *MockStore) CreatePayment(arg0 context.Context, arg1 db.CreatePaymentParams) (db.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", arg0, arg1)
	ret0, _ := ret[0].(db.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockStoreMockRecorder) CreatePayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockStore)(nil).CreatePayment), arg0, arg1)
}

// CreatePizza mocks base method.
func (m *MockStore) CreatePizza(arg0 context.Context, arg1 db.CreatePizzaParams) (db.Pizza, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePizza", arg0, arg1)
	ret0, _ := ret[0].(db.Pizza)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePizza indicates an expected call of CreatePizza.
func (mr *MockStoreMockRecorder) CreatePizza(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePizza", reflect.TypeOf((*MockStore)(nil).CreatePizza), arg0, arg1)
}

// DeleteCustomer mocks base method.
func (m *MockStore) DeleteCustomer(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCustomer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCustomer indicates an expected call of DeleteCustomer.
func (mr *MockStoreMockRecorder) DeleteCustomer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomer", reflect.TypeOf((*MockStore)(nil).DeleteCustomer), arg0, arg1)
}

// GetCustomer mocks base method.
func (m *MockStore) GetCustomer(arg0 context.Context, arg1 int64) (db.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomer", arg0, arg1)
	ret0, _ := ret[0].(db.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomer indicates an expected call of GetCustomer.
func (mr *MockStoreMockRecorder) GetCustomer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockStore)(nil).GetCustomer), arg0, arg1)
}

// GetCustomerForUpdate mocks base method.
func (m *MockStore) GetCustomerForUpdate(arg0 context.Context, arg1 int64) (db.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomerForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomerForUpdate indicates an expected call of GetCustomerForUpdate.
func (mr *MockStoreMockRecorder) GetCustomerForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomerForUpdate", reflect.TypeOf((*MockStore)(nil).GetCustomerForUpdate), arg0, arg1)
}

// GetOrder mocks base method.
func (m *MockStore) GetOrder(arg0 context.Context, arg1 int64) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockStoreMockRecorder) GetOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockStore)(nil).GetOrder), arg0, arg1)
}

// GetPayment mocks base method.
func (m *MockStore) GetPayment(arg0 context.Context, arg1 int64) (db.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayment", arg0, arg1)
	ret0, _ := ret[0].(db.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayment indicates an expected call of GetPayment.
func (mr *MockStoreMockRecorder) GetPayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayment", reflect.TypeOf((*MockStore)(nil).GetPayment), arg0, arg1)
}

// GetPizza mocks base method.
func (m *MockStore) GetPizza(arg0 context.Context, arg1 int64) (db.Pizza, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPizza", arg0, arg1)
	ret0, _ := ret[0].(db.Pizza)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPizza indicates an expected call of GetPizza.
func (mr *MockStoreMockRecorder) GetPizza(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPizza", reflect.TypeOf((*MockStore)(nil).GetPizza), arg0, arg1)
}

// ListCustomers mocks base method.
func (m *MockStore) ListCustomers(arg0 context.Context, arg1 db.ListCustomersParams) ([]db.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCustomers", arg0, arg1)
	ret0, _ := ret[0].([]db.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCustomers indicates an expected call of ListCustomers.
func (mr *MockStoreMockRecorder) ListCustomers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCustomers", reflect.TypeOf((*MockStore)(nil).ListCustomers), arg0, arg1)
}

// ListOrders mocks base method.
func (m *MockStore) ListOrders(arg0 context.Context, arg1 db.ListOrdersParams) ([]db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", arg0, arg1)
	ret0, _ := ret[0].([]db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockStoreMockRecorder) ListOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockStore)(nil).ListOrders), arg0, arg1)
}

// ListPayments mocks base method.
func (m *MockStore) ListPayments(arg0 context.Context, arg1 db.ListPaymentsParams) ([]db.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPayments", arg0, arg1)
	ret0, _ := ret[0].([]db.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPayments indicates an expected call of ListPayments.
func (mr *MockStoreMockRecorder) ListPayments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPayments", reflect.TypeOf((*MockStore)(nil).ListPayments), arg0, arg1)
}

// ListPizzas mocks base method.
func (m *MockStore) ListPizzas(arg0 context.Context, arg1 db.ListPizzasParams) ([]db.Pizza, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPizzas", arg0, arg1)
	ret0, _ := ret[0].([]db.Pizza)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPizzas indicates an expected call of ListPizzas.
func (mr *MockStoreMockRecorder) ListPizzas(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPizzas", reflect.TypeOf((*MockStore)(nil).ListPizzas), arg0, arg1)
}

// UpdateCustomerAddress mocks base method.
func (m *MockStore) UpdateCustomerAddress(arg0 context.Context, arg1 db.UpdateCustomerAddressParams) (db.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomerAddress", arg0, arg1)
	ret0, _ := ret[0].(db.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCustomerAddress indicates an expected call of UpdateCustomerAddress.
func (mr *MockStoreMockRecorder) UpdateCustomerAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomerAddress", reflect.TypeOf((*MockStore)(nil).UpdateCustomerAddress), arg0, arg1)
}

// UpdateOrderStatus mocks base method.
func (m *MockStore) UpdateOrderStatus(arg0 context.Context, arg1 db.UpdateOrderStatusParams) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockStoreMockRecorder) UpdateOrderStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockStore)(nil).UpdateOrderStatus), arg0, arg1)
}

// UpdatePaymentStatus mocks base method.
func (m *MockStore) UpdatePaymentStatus(arg0 context.Context, arg1 db.UpdatePaymentStatusParams) (db.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePaymentStatus", arg0, arg1)
	ret0, _ := ret[0].(db.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePaymentStatus indicates an expected call of UpdatePaymentStatus.
func (mr *MockStoreMockRecorder) UpdatePaymentStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePaymentStatus", reflect.TypeOf((*MockStore)(nil).UpdatePaymentStatus), arg0, arg1)
}
