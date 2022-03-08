package main

import "fmt"

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{builder: b}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setWalls()
	d.builder.setWindows()
	d.builder.setDoors()
	fmt.Println(d.builder)
	return d.builder.getHouse()
}
