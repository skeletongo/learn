package main

import (
	"fmt"
	"reflect"
)

type One struct {
	name *string
}

func f() {
	a := []*int{}
	tp := reflect.TypeOf(&a)
	fmt.Println("1", tp.Name())
	fmt.Println("2", tp.String())
}
