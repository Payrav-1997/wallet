package main

import (
	"github.com/Payrav-1997/wallet/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}
	svc.RegisterAccount("+992915224442")
	svc.FindAccountByID(1)

}
