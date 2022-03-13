package main

import (
	"errors"
	"fmt"
)

type hasStockState struct {
	vendingMachine *vendingMachine
}

func (s *hasStockState) addStock() error {
	fmt.Println("added an item to the stock")
	s.vendingMachine.itemCount += 1
	return nil
}

func (s *hasStockState) requestItem() error {
	if hasItemsInStock(s.vendingMachine.itemCount) {
		fmt.Println("transaction requested")
		s.vendingMachine.setState(s.vendingMachine.transactionRequested)
		return nil
	}

	s.vendingMachine.setState(s.vendingMachine.emptyStock)
	return errors.New("there is no stock, item request failed")
}

func (s *hasStockState) insertMoney(money int) error {
	return errors.New("no item request to pay for")
}

func (s *hasStockState) dispenseItem() error {
	return errors.New("there is no transaction being processed, dispense item failed")
}

func hasItemsInStock(itemQuantity int) bool {
	return itemQuantity > 0
}