package wallet

import (
	
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"bufio"

	"github.com/Payrav-1997/wallet/pkg/types"
	"github.com/google/uuid"
)

var ErrPhoneRegistered = errors.New("phone already registered")
var ErrAmountMustBePositive = errors.New("amount must be greater than zero")
var ErrAccountNotFound = errors.New("account not found")
var ErrNotEnoughBalance = errors.New("not enough balance")
var ErrPaymentNotFound = errors.New("payment not found")
var ErrFavoriteNotFound = errors.New("favorite not found")
var ErrFileNotFound = errors.New("file not found")

type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
	favorites     []*types.Favorite
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
		return nil, ErrNotEnoughBalance
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

func (s *Service) FindPaymentByID(paymentID string) (*types.Payment, error) {
	var payment *types.Payment

	for _, pay := range s.payments {
		if pay.ID == paymentID {
			payment = pay
		}
	}

	if payment == nil {
		return nil, ErrPaymentNotFound
	}

	return payment, nil
}

func (s *Service) FindFavoriteByID(favoriteID string) (*types.Favorite, error) {
	for _, favorite := range s.favorites {
		if favorite.ID == favoriteID {
			return favorite, nil
		}
	}
	return nil, ErrFavoriteNotFound
}

//Deposit method
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

func (s *Service) Reject(paymentID string) error {
	pay, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return err
	}

	acc, err := s.FindAccountByID(pay.AccountID)
	if err != nil {
		return err
	}

	pay.Status = types.PaymentStatusFail
	acc.Balance += pay.Amount

	return nil
}

func (s *Service) Repeat(paymentID string) (*types.Payment, error) {
	pay, err := s.FindPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	payment, err := s.Pay(pay.AccountID, pay.Amount, pay.Category)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *Service) FavoritePayment(paymentID string, name string) (*types.Favorite, error) {
	payment, err := s.FindPaymentByID(paymentID)

	if err != nil {
		return nil, err
	}

	favoriteID := uuid.New().String()
	newFavorite := &types.Favorite{
		ID:        favoriteID,
		AccountID: payment.AccountID,
		Name:      name,
		Amount:    payment.Amount,
		Category:  payment.Category,
	}

	s.favorites = append(s.favorites, newFavorite)
	return newFavorite, nil
}

func (s *Service) PayFromFavorite(favoriteID string) (*types.Payment, error) {
	favorite, err := s.FindFavoriteByID(favoriteID)
	if err != nil {
		return nil, err
	}

	payment, err := s.Pay(favorite.AccountID, favorite.Amount, favorite.Category)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *Service) ExportToFile(path string) error {
	r := ""
	file, err := os.Create(path)
	if err != nil {
		log.Print(err)
		return ErrFileNotFound
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()
	for _, acc := range s.accounts {
		ID := strconv.Itoa(int(acc.ID)) + ";"
		Phone := string(acc.Phone) + ";"
		Balance := strconv.Itoa(int(acc.Balance))

		r += ID
		r += Phone
		r += Balance + "|"
	}
	_, err = file.Write([]byte(r))
	if err != nil {
		log.Print(err)
		return ErrFileNotFound
	}

	return nil
}
func ReadFile(file *os.File) ([]byte, error) {
	content := make([]byte, 0)
	buf := make([]byte, 4)
	for {
		read, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		content = append(content, buf[:read]...)
	}
	return content, nil
}

func AddAccountToFile(file *os.File, account *types.Account) error {
	content, err := ReadFile(file)
	if err != nil {
		return err
	}
	r := string(content)

	ID := strconv.Itoa(int(account.ID))
	Phone := string(account.Phone)
	Balance := strconv.Itoa(int(account.Balance))

	r += (ID + ";")
	r += (Phone + ";")
	r += (Balance + "|")
	file.Write([]byte(r))
	return nil
}

func (s *Service) ImportFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	content, err := ReadFile(file)

	var accounts []string = strings.Split(string(content), "|")

	for _, account := range accounts[:len(accounts)-1] {
		var vals []string = strings.Split(account, ";")
		id, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		balance, err := strconv.Atoi(vals[2])
		if err != nil {
			return err
		}
		newAccount := &types.Account{
			ID:      int64(id),
			Phone:   types.Phone(vals[1]),
			Balance: types.Money(balance),
		}
		s.accounts = append(s.accounts, newAccount)
	}
	return nil
}

func (s *Service) Export(dir string)error{
	if s.accounts != nil{
		file,err := os.Create(dir + "/data/account.dump/")
		if err !=nil{
			log.Print(err)
			return err
		}
			for _,account := range s.accounts{
			acc := []byte(strconv.FormatInt(int64(account.ID),1)+string(";")+string(account.Phone)+string(";") + 
			strconv.FormatInt(int64(account.Balance),1)+	string(";")+ string('\n'))
				_,err = file.Write(acc)
				if err != nil{
					log.Print(err)
					return err
				}
		}
	
	}else {
		log.Print("Аккаунт не найден")
	}
	return nil
}

func (s *Service) Import(dir string)error{

	src,err:= os.Open(dir + "/data/account.dump")
	if err != nil{
		log.Print(err)
	}else{
		defer func(){
			if cerr:= src.Close();
			cerr!=nil{
				log.Print(cerr)
			}
		}()
		reader := bufio.NewReader(src)
		for{
			read,err := reader.ReadString('\n')
			if err == io.EOF{
				log.Print(read)
				break
			}
			if err != nil{
				log.Print(err)
			}
		}
	}
	return nil
}