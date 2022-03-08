package main

type iShoes interface {
	setBrand(brand string)
	setSize(size float32)
}

type shoe struct {
	brand string
	size float32
}

func (s *shoe) setBrand(brand string) {
	s.brand = brand
}

func (s *shoe) setSize(size float32) {
	s.size = size
}

