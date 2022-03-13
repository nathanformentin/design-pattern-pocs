package main

import (
	"errors"
	"fmt"
)

var transactionBeingProcessedErr = errors.New("a transaction is being processed")

type processingTransactionState struct {
	vendingMachine *vendingMachine
}

func (s *processingTransactionState) addStock() error {
	return transactionBeingProcessedErr
}

func (s *processingTransactionState) requestItem() error {
	return transactionBeingProcessedErr
}

func (s *processingTransactionState) insertMoney(money int) error {
	return transactionBeingProcessedErr
}

func (s *processingTransactionState) dispenseItem() error {
	s.vendingMachine.itemCount -= 1
	s.vendingMachine.setState(s.vendingMachine.hasStock)

	fmt.Println("transaction complete")

	return nil
}
