package repository

import (
	"database/sql"
	"ecommerce/model"
	"errors"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (repo *AccountRepository) GetAccountByUsername(username string) (*model.Account, error) {
	row := repo.db.QueryRow(`SELECT id, username, password FROM users WHERE username = ?`, username)
	account := &model.Account{}
	err := row.Scan(&account.ID, &account.UserName, &account.PassKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("account not found")
		}
		return nil, err
	}
	return account, nil
}

func (repo *AccountRepository) RegisterAccount(account *model.Account) error {
	_, err := repo.db.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, account.UserName, account.PassKey)
	return err
}
