package domain

import "fmt"

const CREDIT_CARD_NUMBER_SIZE = 16

// Based in PCI-DSS rules
func SanitizeCreditCardNumber(c string) (string, error) {
	if len(c) != CREDIT_CARD_NUMBER_SIZE {
		return "", NewErrInvalidCreditCardSize(len(c))
	}

	return fmt.Sprintf("%sXXXXXX%s", string(c[:6]), string(c[12:])), nil
}
