package domain

import (
	"go-api-dev/dto"
	"go-api-dev/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
