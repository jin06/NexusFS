package main

import "fmt"

type test interface {
	ret() test
}

type astruct struct {
	a string
}

func (a *astruct) ret() test {
	return a
}

type bstruct struct {
	a string
	as astruct
}

func pp(t test) {
	fmt.Println(t)
}

func main() {
	x := astruct{
		a: "aaa",
	}
	y := bstruct{
		a:       "bbb",
		as: x,
	}
	pp(y.as.ret())
}
