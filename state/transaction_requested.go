package main

import (
	"errors"
	"fmt"
)

var transactionRequestInProgressError = errors.New("transaction request in progress")

type transactionRequestedState struct {
	vendingMachine *vendingMachine
}

func (s *transactionRequestedState) addStock() error {
	return transactionRequestInProgressError
}

func (s *transactionRequestedState) requestItem() error {
	return transactionRequestInProgressError
}

func (s *transactionRequestedState) insertMoney(money int) error {
	if s.isEnoughMoneyToBuyProduct(money) {
		fmt.Println("transaction paid")
		s.vendingMachine.setState(s.vendingMachine.processingTransaction)
		return nil
	}

	return errors.New("not enough money to complete transaction")
}

func (s *transactionRequestedState) dispenseItem() error {
	return transactionRequestInProgressError
}

func (s *transactionRequestedState) isEnoughMoneyToBuyProduct(money int) bool {
	return money > s.vendingMachine.itemPrice
}

