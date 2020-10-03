package main

import (
	"github.com/Payrav-1997/wallet/pkg/wallet"
)

func main() {
	svc := &wallet.Services{}
	svc.RegisterAccount("+992915224442")
	svc.FindAccountByIdmethod(1)
	
}