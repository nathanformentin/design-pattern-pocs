package main

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun{
			name:  "AK47",
			power: 0,
		}}
}

