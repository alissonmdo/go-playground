package accounts

import (
	"bank/clients"
	"fmt"
)

type CheckingAccount struct {
	Number  int
	Owner   clients.Client
	Agency  int
	balance float64
}

func (a *CheckingAccount) Transfer(amount float64, destinationAccount *CheckingAccount) bool {
	if a.balance < amount && amount > 0 {
		return false
	}
	a.balance -= amount
	destinationAccount.Deposit(amount)
	return true
}

func (a *CheckingAccount) Withdraw(amount float64) (string, float64) {
	canWithdraw := amount > 0 && a.balance >= amount
	message := "Insuficient funds"
	if canWithdraw {
		a.balance -= amount
		message = "Withdraw Successful"
	}
	return message, a.balance
}

func (a *CheckingAccount) Deposit(amount float64) (string, float64) {
	canDeposit := amount > 0
	message := "Deposit Failed"
	if canDeposit {
		a.balance += amount
		message = "Deposit successful"
	}
	return message, a.balance
}

func (a *CheckingAccount) Print() {
	fmt.Println("")
	fmt.Println("----------------------")
	fmt.Println("Account Number", a.Number, "- Agency", a.Agency)
	fmt.Println("Owner:", a.Owner)
	fmt.Println("Current balance:", a.balance)
	fmt.Println("----------------------")
}
