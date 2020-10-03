package wallet

import (
	"github.com/Payrav-1997/wallet/pkg/wallet"
	
	"fmt"
	"testing"

)

func TestServices_RegisterAccount(t *testing.T) {
	svc := &wallet.Services{}
	account, err := svc.RegisterAccount("+992915224442")
	if err != nil {
		fmt.Println(account)
	}

	account, err = svc.FindAccountByIdmethod(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestServices_FindAccoundByIdmethod_notFound(t *testing.T) {
	svc := &wallet.Services{}
	account, err := svc.RegisterAccount("+992915224442")
	if err != nil {
		fmt.Println(account)
	}

	account, err = svc.FindAccountByIdmethod(2)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}