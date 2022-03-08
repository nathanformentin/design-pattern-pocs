package main

type nikeFactory struct {

}

func (a nikeFactory) buildShoes() iShoes {
	return nikeShoes{&shoe{
		brand: "adidas",
		size:  15,
	}}
}

func (a nikeFactory) buildShorts() iShorts {
	return nikeShorts{
		&shorts{
			brand: "adidas",
			size:  15.5,
		},
	}
}

