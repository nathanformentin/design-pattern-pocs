package main

type iBuilder interface {
	setWindows()
	setDoors()
	setWalls()
	getHouse() house
}