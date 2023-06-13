package domain

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrCreditCardTransactionRepositoryHouston = errors.New("Houston, we have unknown error into credit card transaction repository")
	ErrTransactionStatusRepositoryHouston     = errors.New("Houston, we have unknown error into transaction status repository")
	ErrAuthorizationServiceHouston            = errors.New("Houston, we have unknown error into authorization service")
)

type ErrPaginationCriteriaPage struct {
	page uint
}

func NewErrPaginationCriteriaPage(page uint) ErrPaginationCriteriaPage {
	return ErrPaginationCriteriaPage{
		page: page,
	}
}

func (e ErrPaginationCriteriaPage) Error() string {
	return fmt.Sprintf("Invalid pagination criteria page [ %d ]", e.page)
}

type ErrPaginationCriteriaItemsPerPage struct {
	itemsPerPage uint
}

func NewErrPaginationCriteriaItemsPerPage(itemsPerPage uint) ErrPaginationCriteriaItemsPerPage {
	return ErrPaginationCriteriaItemsPerPage{
		itemsPerPage: itemsPerPage,
	}
}

func (e ErrPaginationCriteriaItemsPerPage) Error() string {
	return fmt.Sprintf(
		"Invalid pagination criteria items per page, expected between [ %d ] and [ %d ], got [ %d ]",
		ITEMS_PER_PAGE_MIN,
		ITEMS_PER_PAGE_MAX,
		e.itemsPerPage,
	)
}

type ErrCreditCardTransactionNotFound struct {
	criteria string
	value    string
}

func NewErrCreditCardTransactionNotFound(criteria string, value string) ErrCreditCardTransactionNotFound {
	return ErrCreditCardTransactionNotFound{
		criteria: criteria,
		value:    value,
	}
}

func (e ErrCreditCardTransactionNotFound) Error() string {
	return fmt.Sprintf("Credit card transaction not found by criteria [ %s ] and value [ %s ]", e.criteria, e.value)
}

type ErrInvalidCreditCardNumber struct {
	number string
}

func NewErrInvalidCreditCardNumber(number string) ErrInvalidCreditCardNumber {
	return ErrInvalidCreditCardNumber{
		number: number,
	}
}

func (e ErrInvalidCreditCardNumber) Error() string {
	return fmt.Sprintf("Invalid credit card number [ %s ]", e.number)
}

type ErrCreditCardTransactionInstallments struct {
	installments uint
}

func NewErrCreditCardTransactionInstallments(installments uint) ErrCreditCardTransactionInstallments {
	return ErrCreditCardTransactionInstallments{
		installments: installments,
	}
}

func (e ErrCreditCardTransactionInstallments) Error() string {
	return fmt.Sprintf(
		"Invalid installments, expected between [ %d ] and [ %d ], got [ %d ]",
		INSTALLMENTS_MIN,
		INSTALLMENTS_MAX,
		e.installments,
	)
}

type ErrTransactionStatusInvalid struct {
	status string
}

func NewErrTransactionStatusInvalid(status string) ErrTransactionStatusInvalid {
	return ErrTransactionStatusInvalid{
		status: status,
	}
}

func (e ErrTransactionStatusInvalid) Error() string {
	return fmt.Sprintf(
		"Invalid transaction status, expected one of [ %s ], got [ %s ]",
		strings.Join(GetTransactionStatusAvailable(), ", "),
		e.status,
	)
}

type ErrAuthorizationResponseStatusInvalid struct {
	status string
}

func NewErrAuthorizationResponseStatusInvalid(status string) ErrAuthorizationResponseStatusInvalid {
	return ErrAuthorizationResponseStatusInvalid{
		status: status,
	}
}

func (e ErrAuthorizationResponseStatusInvalid) Error() string {
	return fmt.Sprintf(
		"Invalid authorization response status, expected one of [ %s ], got [ %s ]",
		strings.Join(GetAuthorizationStatusAvailable(), ", "),
		e.status,
	)
}
