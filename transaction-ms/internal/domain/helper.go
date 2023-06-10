package domain

import (
	"fmt"
	"regexp"

	"github.com/oklog/ulid/v2"
)

const (
	CREDIT_CARD_NUMBER_REGEX = `^[\d]{16}$`
)

func IsValidCardNumber(c string) bool {
	return regexp.MustCompile(CREDIT_CARD_NUMBER_REGEX).MatchString(c)
}

// Based in PCI-DSS rules
func SanitizeCreditCardNumber(c string) (string, error) {
	if !IsValidCardNumber(c) {
		return "", NewErrInvalidCreditCardNumber(c)
	}

	return fmt.Sprintf("%sXXXXXX%s", string(c[:6]), string(c[12:])), nil
}

func TransactionID() string {
	return ulid.Make().String()
}
