package infra

type CreateCreditCardTransactionRequest struct {
	HolderName   string `json:"holder_name" binding:"required"`
	CardNumber   string `json:"card_number" binding:"required,min=16,max=16"`
	CVV          string `json:"cvv" binding:"required,min=3,max=3"`
	ExpireYear   string `json:"expire_year" binding:"required,min=4,max=4"`
	ExpireMonth  string `json:"expire_month" binding:"required,min=2,max=2"`
	Amount       uint   `json:"amount" binding:"required"`
	Installments uint   `json:"installments" binding:"required"`
	Description  string `json:"description" binding:"required"`
}

type EditCreditCardTransactionRequest struct {
	Description string `json:"description" binding:"required"`
}
