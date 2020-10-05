package wallet

import (	
	"testing"
)

func TestService_RegisterAccount_success(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+992915224442")

	account, err := svc.FindAccountByID(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_FindAccoundById_notFound(t *testing.T) {
	SVC := Service{}
	SVC.RegisterAccount("+992915224442")

	account, err := SVC.FindAccountByID(2)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_Reject_success(t *testing.T){
	SVC:= Service{}
	SVC.RegisterAccount("+992915224442")

	account,err := SVC.FindAccountByID(1)
	if err != nil{
		t.Errorf("\ngot> %v \nwant > nil",err)
	}
	err = SVC.Deposit(account.ID, 1000_00)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	payment, err := SVC.Pay(account.ID, 100_00, "Caffe")
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

	payment, err := SVC.Pay(account.ID, 100_00, "caffe")
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	pay, err := SVC.FindPaymentByID(payment.ID)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}

	editPayID := pay.ID + "caffe"
	err = SVC.Reject(editPayID)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}
