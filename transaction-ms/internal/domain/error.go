package domain

import (
	"fmt"
	"strings"
)

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

type ErrCreditCardTransctionInstallments struct {
	installments uint
}

func NewErrCreditCardTransctionInstallments(installments uint) ErrCreditCardTransctionInstallments {
	return ErrCreditCardTransctionInstallments{
		installments: installments,
	}
}

func (e ErrCreditCardTransctionInstallments) Error() string {
	return fmt.Sprintf(
		"Invalid installments, expected between [ %d ] and [ %d ], got [ %d ]",
		INSTALLMENTS_MIN,
		INSTALLMENTS_MAX,
		e.installments,
	)
}

type ErrTransctionStatusInvalid struct {
	status string
}

func NewErrTransctionStatusInvalid(status string) ErrTransctionStatusInvalid {
	return ErrTransctionStatusInvalid{
		status: status,
	}
}

func (e ErrTransctionStatusInvalid) Error() string {
	return fmt.Sprintf(
		"Invalid transction status, expected one of [ %s ], got [ %s ]",
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
