package main

import (
	"github.com/andodeki/code/HA/api.gridbackendapp.com/test/here"
)

type A struct {
	b here.B
}

func (a *A) updateB(n int) {
	a.b.C = n
}

func main() {
	// a := A{b: B{C:5}}

	// fmt.Println(a)
	// a.updateB(42)
	// fmt.Println(a)
}
