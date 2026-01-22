package main

import (
	"errors"
	"fmt"
)

var ErrNegativeAmount = errors.New("negative withdrawal amount")
var ErrInsufficientBalance = errors.New("insufficient balance")

func withdraw(balance int, amount int) (int, error) {
	if amount < 0 {
		return balance, ErrNegativeAmount
	}
	if amount > balance {
		return balance, ErrInsufficientBalance
	}
	return balance - amount, nil
}

func withDraw() {
	balance := 1000

	amounts := []int{200, -50, 5000, 300}

	for _, amt := range amounts {
		newBalance, err := withdraw(balance, amt)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		balance = newBalance
		fmt.Println("Withdrawal successful. New balance:", balance)
	}
}
