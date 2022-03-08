package main

import "errors"

type sportsFactory interface {
	buildShoes() iShoes
	buildShorts() iShorts
}

func getSportsFactory(brand string) (sportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidasFactory{}, nil
	case "nike":
		return &nikeFactory{}, nil
	}

	return nil, errors.New("invalid brand")
}
