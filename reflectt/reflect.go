package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type InnerType struct {
	//Bool        bool
	//Int         int
	//Int8        int8
	//Int16       int16
	//Int32       int32
	//Int64       int64
	//Uint        uint
	//Uint8       uint8
	//Uint16      uint16
	//Uint32      uint32
	//Uint64      uint64
	//Uintptr     uintptr
	//Float32     float32
	//Float64     float64
	//Complex64   complex64
	//Complex128  complex128
	Func        func()
	Chan        chan int
	Ptr         *int
	Struct      Student
	Array       [2]int
	Slice       []int
	Interface   Person
	Map         map[string]interface{}
	String      string
	UnsafePoint unsafe.Pointer
}

type Student struct {
}

func (s *Student) Speak(str string) {
	fmt.Println(str)
}

type Person interface {
	Speak(string)
}

func main() {
	testFunc()
}

func testFunc() {
	data := InnerType{
		Func: func() {
			fmt.Println("hello world!")
		},
	}
	t := reflect.TypeOf(data.Func)
	fmt.Println("Name()=", t.Name())
	fmt.Println("Kind()=", t.Kind())
	fmt.Println("String()=", t.String())
	fmt.Println("Align()=", t.Align())
	fmt.Println("Size()=", t.Size())
	v := reflect.ValueOf(data.Func)
	fmt.Println("Kind()=", v.Kind())
	fmt.Println("String()=", v.String())
}
