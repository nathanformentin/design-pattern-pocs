package main

type iShorts interface {
	setBrand(brand string)
	setSize(size float32)
}
type shorts struct {
	brand string
	size float32
}

func (s *shorts) setBrand(brand string) {
	s.brand = brand
}

func (s *shorts) setSize(size float32) {
	s.size = size
}

