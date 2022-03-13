package main

import "fmt"

type folder struct {
	internalInodes []inode
	name string
}

func (f folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, inode := range f.internalInodes {
		inode.print(indentation + indentation)
	}
}

func (f folder) clone() inode {
	return &folder{
		internalInodes: f.internalInodes,
		name:           f.name,
	}
}

