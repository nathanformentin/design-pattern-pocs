package main

type adidasFactory struct {
	
}

func (a adidasFactory) buildShoes() iShoes {
	return adidasShoes{&shoe{
		brand: "adidas",
		size:  15,
	}}
}

func (a adidasFactory) buildShorts() iShorts {
	return adidasShorts{
		&shorts{
			brand: "adidas",
			size:  15.5,
		},
	}
}

