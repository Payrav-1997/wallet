package wallet

import (	
	"testing"
)

func TestService_RegisterAccount_success(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+992915224442")

	account, err := svc.FindAccountBy(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_FindAccoundById_notFound(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("+992915224442")

	account, err := svc.FindAccountById(2)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}