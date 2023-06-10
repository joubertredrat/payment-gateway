package domain

import "time"

const (
	INSTALLMENTS_MIN  = 1
	INSTALLMENTS_MAX  = 12
	STATUS_CREATED    = "created"
	STATUS_AUTHORIZED = "authorized"
	STATUS_REFUSED    = "refused"
	STATUS_FINISHED   = "finished"
)

type CreditCardTransction struct {
	ID               string
	CardNumber       string
	Amount           uint
	Installments     uint
	Description      string
	TransctionStatus []TransctionStatus
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}

type TransctionStatus struct {
	ID                     string
	CreditCardTransctionID string
	Status                 string
	CreatedAt              *time.Time
}

func IsValidInstallments(i uint) bool {
	return i >= INSTALLMENTS_MIN && i <= INSTALLMENTS_MAX
}

func GetStatusAvailable() []string {
	return []string{STATUS_CREATED, STATUS_AUTHORIZED, STATUS_REFUSED, STATUS_FINISHED}
}

func IsValidStatus(v string) bool {
	return contains(v, GetStatusAvailable())
}

func contains(v string, e []string) bool {
	for _, s := range e {
		if v == s {
			return true
		}
	}
	return false
}
