package domain

import "fmt"

type ErrInvalidCreditCardSize struct {
	size int
}

func NewErrInvalidCreditCardSize(size int) ErrInvalidCreditCardSize {
	return ErrInvalidCreditCardSize{
		size: size,
	}
}

func (e ErrInvalidCreditCardSize) Error() string {
	return fmt.Sprintf(
		"Invalid credit card number size, expected [ %d ], got [ %d ]",
		CREDIT_CARD_NUMBER_SIZE,
		e.size,
	)
}
