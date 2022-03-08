package main

import "errors"

func getHouseBuilder(builderType string) (iBuilder, error) {
	switch builderType {
	case "normal":
		return &newNormalHouseBuilder(), nil
	case "igloo":
		return &newIglooBuilder(), nil
	}

	return nil, errors.New("invalid builder type")
}
