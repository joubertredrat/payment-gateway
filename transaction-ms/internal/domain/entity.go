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
	ID           string
	CardNumber   string
	Amount       uint
	Installments uint
	Description  string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type TransctionStatus struct {
	Status    string
	CreatedAt *time.Time
}
