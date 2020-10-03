package wallet

import (
	"errors"
	"github.com/Payrav-1997/wallet/pkg/types"
)

var (
	ErrPhoneRegistered      = errors.New("телефон уже зарегистрирован")         
	ErrAmountMustBePositive = errors.New("сумма должна быть больше нуля") 
	ErrAccountNotFound      = errors.New("аккаунт не найден")                
)


type Services struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

func (s *Services) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

func (s *Services) FindAccountByIdmethod(accountID int64) (*types.Account, error) {
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}

	if account == nil {
		return nil, ErrAccountNotFound
	}

	return account, nil
}