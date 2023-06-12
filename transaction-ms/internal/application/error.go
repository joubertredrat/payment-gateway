package application

import "errors"

var (
	ErrUsecaseCreateCreditCardTransactionHouston = errors.New("Houston, we have unknown error into create credit card transaction use case")
	ErrUsecaseEditCreditCardTransactionHouston   = errors.New("Houston, we have unknown error into edit credit card transaction use case")
	ErrUsecaseDeleteCreditCardTransactionHouston = errors.New("Houston, we have unknown error into delete credit card transaction use case")
	ErrUsecaseGetCreditCardTransactionHouston    = errors.New("Houston, we have unknown error into get credit card transaction use case")
	ErrDispatcherHouston                         = errors.New("Houston, we have unknown error into dispatcher")
)
