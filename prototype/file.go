package main

import "fmt"

const _cloneSuffix = "_clone"

type file struct {
	name string
}

func (f file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f file) clone() inode {
	return &file{name: f.name + _cloneSuffix}
}



