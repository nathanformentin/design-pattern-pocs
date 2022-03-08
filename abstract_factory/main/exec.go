package main

import "fmt"

func main() {
	nikeFactory, _ := getSportsFactory("nike")
	adidasFactory, _ := getSportsFactory("adidas")

	shoes := nikeFactory.buildShoes()
	shorts := nikeFactory.buildShorts()
	adidasFactory.buildShoes()
	fmt.Println(fmt.Sprintf("shoe brand %s, shoe size %d", shoes)
	fmt.Println(shorts)
}
