package infra

type RequestValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type ListResponse struct {
	Metadata MetadataResponse `json:"metadata"`
	Data     interface{}      `json:"data"`
}

func NewListResponse(page, itemsPerPage uint, data interface{}) ListResponse {
	return ListResponse{
		Metadata: MetadataResponse{
			Pagination: PaginationResponse{
				Page:         page,
				ItemsPerPage: itemsPerPage,
			},
		},
		Data: data,
	}
}

type MetadataResponse struct {
	Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	Page         uint `json:"page"`
	ItemsPerPage uint `json:"items_per_page"`
}

type AmountResponse struct {
	Cents uint   `json:"cents"`
	BRL   string `json:"brl"`
}

type CreditCardTransactionResponse struct {
	TransactionID     string                      `json:"transaction_id"`
	CardNumber        string                      `json:"card_number"`
	Amount            AmountResponse              `json:"amount"`
	Installments      uint                        `json:"installments"`
	Description       string                      `json:"description"`
	TransactionStatus []TransactionStatusResponse `json:"transaction_status"`
	CreatedAt         *string                     `json:"created_at"`
	UpdatedAt         *string                     `json:"updated_at"`
}

type CreditCardTransactionListResponse struct {
	TransactionID string         `json:"transaction_id"`
	CardNumber    string         `json:"card_number"`
	Amount        AmountResponse `json:"amount"`
	Installments  uint           `json:"installments"`
	Description   string         `json:"description"`
	CreatedAt     *string        `json:"created_at"`
	UpdatedAt     *string        `json:"updated_at"`
}

type TransactionStatusResponse struct {
	Status    string  `json:"status"`
	CreatedAt *string `json:"created_at"`
}
