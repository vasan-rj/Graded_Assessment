package service

import (
	"ecommerce/model"
	"ecommerce/repository"
	"errors"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

// RegisterAccount registers a new account.
func (service *AccountService) RegisterAccount(account *model.Account) error {
	// Validate account data
	if account.UserName == "" || account.PassKey == "" {
		return errors.New("invalid account data")
	}

	// Register the account in the database
	return service.repo.RegisterAccount(account)
}

// ValidateCredentials checks if the account credentials are valid.
func (service *AccountService) ValidateCredentials(username, passkey string) (*model.Account, error) {
	account, err := service.repo.GetAccountByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if account.PassKey != passkey {
		return nil, errors.New("incorrect passkey")
	}

	return account, nil
}
