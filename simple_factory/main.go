package main

import (
	"fmt"
)

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	detailsMessage := fmt.Sprintf("gun name %s, power %d", g.getName(), g.getPower())
	fmt.Println(detailsMessage)
}
