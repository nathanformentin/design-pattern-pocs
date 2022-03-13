package main

type state interface {
	addStock() error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
