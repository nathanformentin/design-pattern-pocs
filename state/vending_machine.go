package main

type vendingMachine struct {
	hasStock              state
	transactionRequested  state
	processingTransaction state
	emptyStock            state

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount int, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	v.hasStock = &hasStockState{v}
	v.emptyStock = &emptyStockState{v}
	v.transactionRequested = &transactionRequestedState{v}
	v.processingTransaction = &processingTransactionState{v}

	v.setState(v.hasStock)

	return v
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) addStock() error {
	return v.currentState.addStock()
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}
