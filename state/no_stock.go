package main

import (
	"errors"
)

type emptyStockState struct {
	vendingMachine *vendingMachine
}

func (n *emptyStockState) addStock() error {
	n.vendingMachine.itemCount += 1

	return nil
}

func (n *emptyStockState) requestItem() error {
	return errors.New("no stock! request item failed")
}

func (n *emptyStockState) insertMoney(money int) error {
	return errors.New("no item request to pay for and there is no items in stock")
}

func (n *emptyStockState) dispenseItem() error {
	return errors.New("no items to dispense")
}

