package main

import "errors"

func getGun(gun string) (iGun, error) {
	switch gun {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	}

	return nil, errors.New("invalid gun name")
}
