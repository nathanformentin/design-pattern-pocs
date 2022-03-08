package main

import "fmt"

func main() {
	normalHouseBuilder := newNormalHouseBuilder()

	iglooHouseBuilder := newIglooBuilder()

	director := newDirector(normalHouseBuilder)

	house := director.buildHouse()

	fmt.Println(formatHouseInformation(house))

	director.setBuilder(&iglooHouseBuilder)
	igloo := director.buildHouse()
	fmt.Println(formatHouseInformation(igloo))
}

func formatHouseInformation(h house) string {
	return fmt.Sprintf("window type %s, door type %s, wall type %s", h.windowType, h.doorType, h.wallType)
}


