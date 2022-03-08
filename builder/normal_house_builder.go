package main

type normalHouseBuilder struct {
	windowType string
	doorType string
	wallType string
}

func newNormalHouseBuilder() *normalHouseBuilder {
	return &normalHouseBuilder{}
}

func (b normalHouseBuilder) getHouse() house {
	return house{
		windowType: b.windowType,
		doorType:   b.doorType,
		wallType:   b.wallType,
	}
}

func (b *normalHouseBuilder) setWindows() {
	b.windowType = "normal window"
}

func (b *normalHouseBuilder) setDoors() {
	b.doorType = "normal door"
}

func (b *normalHouseBuilder) setWalls() {
	b.wallType = "normal wall"
}

