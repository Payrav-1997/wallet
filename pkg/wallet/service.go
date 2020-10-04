package wallet

import (
	"errors"
	"github.com/Payrav-1997/wallet/pkg/types"
)

var (
	ErrPhoneRegistered      = errors.New("телефон уже зарегистрирован")         
	ErrAmountMustBePositive = errors.New("сумма должна быть больше нуля") 
	ErrAccountNotFound      = errors.New("аккаунт не найден")
	ErrPaymentNotFound = errors.New("Платеж не найден")                
)

func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {
	if amount <= 0 {
		return nil, ErrAmountMustBePositive
	}

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

	if account.Balance < amount {
		return nil, ErrNotEnoughtBalance
	}

	account.Balance -= amount

	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID:        paymentID,
		AccountID: accountID,
		Amount:    amount,
		Category:  category,
		Status:    types.PaymentStatusInProgress,
	}

	s.payments = append(s.payments, payment)
	return payment, nil
}

type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
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

func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
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


func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount < 0 {
		return ErrAmountMustBePositive
	}

	account, err := s.FindAccountByID(accountID)
	if err != nil {
		return err
	}

	account.Balance += amount
	return nil
}



func (s *Service) FindPaymentByID(paymentID string)(*types.Payment,error){
	var payment *types.Payment
	for _, pay_ment := range s.payments {
		if pay_ment.ID== paymentID{
			payment =pay_ment
		}
	}
	if payment == nil{
		return nil,ErrPaymentNotFound
	}
	return payment,nil
}

func (s *Service) Reject(paymentID string) error{
	pay_ment,err:=s.FindPaymentByID(paymentID)
	if err != nil{
		return err
	}

	acc, err := s.FindAccountByID(pay_ment.ID)
	if err != nil{
		return err
	}
	pay_ment.Status = types.PaymentStatusFail
	acc.Balance += pay_ment.Amount
}