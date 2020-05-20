package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println(f1())
	fmt.Println(f2())
}

func f1() string {
	var s strings.Builder
	s.WriteString(fmt.Sprint(os.Args[0]))
	for _,v:= range os.Args[1:] {
		s.WriteString(fmt.Sprint(" ",v))
	}
	return s.String()
}

func f2() string {
	return strings.Join(os.Args," ")
}