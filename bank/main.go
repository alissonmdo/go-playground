package main

import (
	"bank/accounts"
	"bank/clients"
)

func main() {

	firstAccount := accounts.CheckingAccount{
		Number: 345678,
		Owner: clients.Client{
			Name:             "Alisson de Oliveira",
			RegistrationCode: "123",
			Occupation:       "Software Engineer"},
		Agency: 1}
	firstAccount.Deposit(500)
	secondAccount := accounts.CheckingAccount{Number: 123456, Owner: clients.Client{Name: "Nossila ed Arievilo"}, Agency: 1}
	firstAccount.Transfer(1250, &secondAccount)
	firstAccount.Print()
	secondAccount.Print()
}
