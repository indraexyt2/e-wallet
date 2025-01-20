package models

import "github.com/go-playground/validator/v10"

type TransactionRequest struct {
	Reference string  `json:"reference" validate:"required"`
	Amount    float64 `json:"amount" validate:"required"`
}

func (I *TransactionRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

type ExternalTransactionRequest struct {
	Amount          float64 `json:"amount" validate:"required"`
	Reference       string  `json:"reference" validate:"required"`
	TransactionType string  `json:"transaction_type" validate:"required"`
	WalletID        int     `json:"wallet_id" validate:"required"`
}

func (I *ExternalTransactionRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}
