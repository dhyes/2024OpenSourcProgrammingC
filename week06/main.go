package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 13
	f := 12.9
	c1 := 'Z'
	c2 := '김'
	// fmt.Printf("%d* %f=%f\n", i, f, float64(i)*f)
	fmt.Printf("%d* %f=%d\n", i, f, i*int(f))
	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))
}
