package main

type Struct struct {
	name string
}

func NewStruct() Struct {
	return Struct{
		name: "Foo",
	}
}
