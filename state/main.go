package main

import "log"

func main() {
	vendingMachine := newVendingMachine(10, 1)

	if err := vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vendingMachine.insertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}

	if err := vendingMachine.dispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}
}
