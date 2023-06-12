// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/dispatcher.go

// Package pkg is a generated GoMock package.
package pkg

import (
	domain "joubertredrat/transaction-ms/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDispatcher is a mock of Dispatcher interface.
type MockDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherMockRecorder
}

// MockDispatcherMockRecorder is the mock recorder for MockDispatcher.
type MockDispatcherMockRecorder struct {
	mock *MockDispatcher
}

// NewMockDispatcher creates a new mock instance.
func NewMockDispatcher(ctrl *gomock.Controller) *MockDispatcher {
	mock := &MockDispatcher{ctrl: ctrl}
	mock.recorder = &MockDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDispatcher) EXPECT() *MockDispatcherMockRecorder {
	return m.recorder
}

// CreditCardTransactionCreated mocks base method.
func (m *MockDispatcher) CreditCardTransactionCreated(arg0 domain.CreditCardTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditCardTransactionCreated", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreditCardTransactionCreated indicates an expected call of CreditCardTransactionCreated.
func (mr *MockDispatcherMockRecorder) CreditCardTransactionCreated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditCardTransactionCreated", reflect.TypeOf((*MockDispatcher)(nil).CreditCardTransactionCreated), arg0)
}

// CreditCardTransactionDeleted mocks base method.
func (m *MockDispatcher) CreditCardTransactionDeleted(TransactionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditCardTransactionDeleted", TransactionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreditCardTransactionDeleted indicates an expected call of CreditCardTransactionDeleted.
func (mr *MockDispatcherMockRecorder) CreditCardTransactionDeleted(TransactionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditCardTransactionDeleted", reflect.TypeOf((*MockDispatcher)(nil).CreditCardTransactionDeleted), TransactionID)
}

// CreditCardTransactionEdited mocks base method.
func (m *MockDispatcher) CreditCardTransactionEdited(arg0 domain.CreditCardTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditCardTransactionEdited", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreditCardTransactionEdited indicates an expected call of CreditCardTransactionEdited.
func (mr *MockDispatcherMockRecorder) CreditCardTransactionEdited(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditCardTransactionEdited", reflect.TypeOf((*MockDispatcher)(nil).CreditCardTransactionEdited), arg0)
}

// CreditCardTransactionGot mocks base method.
func (m *MockDispatcher) CreditCardTransactionGot(arg0 domain.CreditCardTransaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditCardTransactionGot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreditCardTransactionGot indicates an expected call of CreditCardTransactionGot.
func (mr *MockDispatcherMockRecorder) CreditCardTransactionGot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditCardTransactionGot", reflect.TypeOf((*MockDispatcher)(nil).CreditCardTransactionGot), arg0)
}
