package main

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun{
		name:  "Musket",
		power: 5,
	}}
}

