package main

import "fmt"

type ATM struct {
	balance float64
}

func (a *ATM) deposit(amount float64) {
	a.balance += amount
}

func (a *ATM) withdraw(amount float64) bool {
	if a.balance < amount {
		return false
	}
	a.balance -= amount
	return true
}

func main() {
	atm := ATM{balance: 0}
	var choice int
	var amount float64

	for {
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is \n ", atm.balance)

		case 2:
			fmt.Println("How much money would you like to deposit?")
			fmt.Scan(&amount)
			atm.deposit(amount)
			fmt.Println("Deposit successful")

		case 3:
			fmt.Println("How much money would you like to withdraw?")
			fmt.Scan(&amount)
			if amount > atm.balance {
				fmt.Println("You don't have enough money")
			}
			atm.withdraw(amount)

		case 4:
			fmt.Print("Exiting...")
			return
		}

	}
}
