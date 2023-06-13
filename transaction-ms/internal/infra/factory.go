package infra

import "joubertredrat/transaction-ms/internal/domain"

func CreateCreditCardTransactionResponseFromUsecase(c domain.CreditCardTransaction) CreditCardTransactionResponse {
	sl := []TransactionStatusResponse{}

	for _, s := range c.TransactionStatus {
		sl = append(sl, CreateTransactionStatusResponseFromUsecase(s))
	}

	return CreditCardTransactionResponse{
		TransactionID: c.TransactionID,
		CardNumber:    c.CardNumber,
		Amount: AmountResponse{
			Cents: c.Amount.Value,
			BRL:   c.Amount.BRL(),
		},
		Installments:      c.Installments,
		Description:       c.Description,
		TransactionStatus: sl,
		CreatedAt:         DatetimeCanonical(c.CreatedAt),
		UpdatedAt:         DatetimeCanonical(c.UpdatedAt),
	}
}

func CreateTransactionStatusResponseFromUsecase(s domain.TransactionStatus) TransactionStatusResponse {
	return TransactionStatusResponse{
		Status:    s.Status,
		CreatedAt: DatetimeCanonical(s.CreatedAt),
	}
}

func CreateCreditCardTransactionListResponseFromUsecase(l []domain.CreditCardTransaction) []CreditCardTransactionListResponse {
	rl := []CreditCardTransactionListResponse{}

	for _, c := range l {
		rl = append(rl, CreditCardTransactionListResponse{
			TransactionID: c.TransactionID,
			CardNumber:    c.CardNumber,
			Amount: AmountResponse{
				Cents: c.Amount.Value,
				BRL:   c.Amount.BRL(),
			},
			Installments: c.Installments,
			Description:  c.Description,
			CreatedAt:    DatetimeCanonical(c.CreatedAt),
			UpdatedAt:    DatetimeCanonical(c.UpdatedAt),
		})
	}

	return rl
}
