package wallet

import (	
	"testing"
)

func TestService_RegisterAccount_success(t *testing.T) {
	svc := Services{}
	svc.RegisterAccount("+992915224442")

	account, err := svc.FindAccountByIdmethod(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}

func TestService_FindAccoundByIdmethod_notFound(t *testing.T) {
	svc := Services{}
	svc.RegisterAccount("+992915224442")

	account, err := svc.FindAccountByIdmethod(2)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", account)
	}
}