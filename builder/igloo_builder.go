package main

type iglooBuilder struct {
	windowType string
	doorType string
	wallType string
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b iglooBuilder) getHouse() house{
	return house{
		windowType: b.windowType,
		doorType: b.doorType,
		wallType: b.wallType,
	}
}

func (b iglooBuilder) setWindows() {
	b.windowType = "igloo window"
}

func (b iglooBuilder) setDoors() {
	b.doorType = "igloo door"
}

func (b iglooBuilder) setWalls() {
	b.wallType = "ice walls"
}