package main

import (
	"errors"
	"fmt"
)

type UserAccount struct {
	AccountID           string
	HolderName          string
	CurrentBalance      float64
	ActivityLog         []string
}

var userAccounts []UserAccount

const (
	OptDeposit            = 1
	OptWithdraw           = 2
	OptViewBalance        = 3
	OptViewTransactions   = 4
	OptExit               = 5
)

func createAccount(accID, holderName string, startingBalance float64) {
	userAccounts = append(userAccounts, UserAccount{
		AccountID:      accID,
		HolderName:     holderName,
		CurrentBalance: startingBalance,
		ActivityLog:    []string{"Account opened with balance: $" + fmt.Sprintf("%.2f", startingBalance)},
	})
}

func addFunds(accID string, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must exceed zero")
	}
	for i, acc := range userAccounts {
		if acc.AccountID == accID {
			userAccounts[i].CurrentBalance += amount
			userAccounts[i].ActivityLog = append(userAccounts[i].ActivityLog, fmt.Sprintf("Deposited: $%.2f", amount))
			return nil
		}
	}
	return errors.New("account not located")
}

func removeFunds(accID string, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must exceed zero")
	}
	for i, acc := range userAccounts {
		if acc.AccountID == accID {
			if acc.CurrentBalance < amount {
				return errors.New("insufficient balance")
			}
			userAccounts[i].CurrentBalance -= amount
			userAccounts[i].ActivityLog = append(userAccounts[i].ActivityLog, fmt.Sprintf("Withdrew: $%.2f", amount))
			return nil
		}
	}
	return errors.New("account not located")
}

func checkBalance(accID string) (float64, error) {
	for _, acc := range userAccounts {
		if acc.AccountID == accID {
			return acc.CurrentBalance, nil
		}
	}
	return 0, errors.New("account not located")
}

func showTransactionLog(accID string) error {
	for _, acc := range userAccounts {
		if acc.AccountID == accID {
			fmt.Println("Transaction Activity:")
			for _, log := range acc.ActivityLog {
				fmt.Println(log)
			}
			return nil
		}
	}
	return errors.New("account not located")
}

func main() {
	createAccount("1001", "Karthik", 2000)
	createAccount("1002", "Meenakshi", 750)

	var option int
	for {
		fmt.Println("\nDigital Banking System")
		fmt.Println("1. Add Funds")
		fmt.Println("2. Withdraw Funds")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Transaction Log")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&option)

		var accID string
		var amount float64

		switch option {
		case OptDeposit:
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&accID)
			fmt.Print("Enter Deposit Amount: ")
			fmt.Scanln(&amount)
			if err := addFunds(accID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}

		case OptWithdraw:
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&accID)
			fmt.Print("Enter Withdrawal Amount: ")
			fmt.Scanln(&amount)
			if err := removeFunds(accID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}

		case OptViewBalance:
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&accID)
			if balance, err := checkBalance(accID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Account Balance: $%.2f\n", balance)
			}

		case OptViewTransactions:
			fmt.Print("Enter Account ID: ")
			fmt.Scanln(&accID)
			if err := showTransactionLog(accID); err != nil {
				fmt.Println("Error:", err)
			}

		case OptExit:
			fmt.Println("Exiting system. Thank you!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
