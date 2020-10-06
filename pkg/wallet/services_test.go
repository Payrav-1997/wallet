package wallet

import (	
	"testing"
)

func TestService_RegisterAccount_success(t *testing.T) {
	SVC := Service{}
	SVC.RegisterAccount("+992915224442")

	account, err := SVC.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_FindAccoundByIdMetod_notFound(t *testing.T) {
	SVC := Service{}
	SVC.RegisterAccount("+992915224442")
	SVC.RegisterAccount("+992915224442")

	account, err := SVC.FindAccountByID(2)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestDeposit(t *testing.T) {
	SVC := Service{}

	SVC.RegisterAccount("+992915224442")

	err := SVC.Deposit(1, 100_00)
	if err != nil {
		t.Error("Произошла ошибка при оплате")
	}

	account, err := SVC.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_Reject_success(t *testing.T) {
	SVC := Service{}
	SVC.RegisterAccount("+992915224442")

	account, err := SVC.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = SVC.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := SVC.Pay(account.ID, 100_00, "Cafe")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := SVC.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = SVC.Reject(pay.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestService_Reject_fail(t *testing.T) {
	SVC := Service{}
	SVC.RegisterAccount("+992915224442")

	account, err := SVC.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = SVC.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := SVC.Pay(account.ID, 100_00, "Cafe")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := SVC.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	editPayID := pay.ID + "mr.virus :)"
	err = SVC.Reject(editPayID)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestService_Repeat_succes(t *testing.T) {
	SVC := Service{}
	SVC.RegisterAccount("+992915224442")

	account, err := SVC.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	err = SVC.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := SVC.Pay(account.ID, 100_00, "Cafe")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := SVC.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err = SVC.Repeat(pay.ID)
	if err != nil {
		t.Errorf("Repeat(): Error(): can't pay (%v): %v", pay.ID, err)
	}
}


func TestService_Favorite_succes(t *testing.T) {
	SVC := Service{}

	account, err := SVC.RegisterAccount("+992915224442")
	if err != nil {
		t.Errorf("RegisterAccount не возвратил ошибку nil , account => %v", account)
	}

	err = SVC.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("method Deposit не возвратил ошибку nil, error => %v", err)
	}

	payment, err := SVC.Pay(account.ID, 10_00, "auto")
	if err != nil {
		t.Errorf("Pay() Error() не могу платить за счет(%v): %v", account, err)
	}

	favorite, err := SVC.FavoritePayment(payment.ID, "megafon")
	if err != nil {
		t.Errorf("FavoritePayment() Error() не могу для favorite(%v): %v", favorite, err)
	}

	paymentFavorite, err := SVC.PayFromFavorite(favorite.ID)
	if err != nil {
		t.Errorf("PayFromFavorite() Error() не могу для favorite(%v): %v", paymentFavorite, err)
	}
}