package main

import (
	"example/bank/fileops"
	"fmt"
)

const balanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(balanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Print("--------------------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

		var option int
		fmt.Print("Your choice: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Enter the amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Your new balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFile)
		case 3:
			fmt.Print("Enter the amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Insufficient balance.")
				continue
			}

			accountBalance -= withdrawAmount // accountBalance = accountBalance - withdrawAmount
			fmt.Println("Your new balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go Bank!")
			return
		}
	}

}
