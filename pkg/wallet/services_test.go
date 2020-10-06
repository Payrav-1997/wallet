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
