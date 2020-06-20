package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type SS struct {
	Name string
	age  int
}

var (
	Bool                   = true
	Int        int         = 1
	Int8       int8        = 1
	Int16      int16       = 1
	Int32      int32       = 1
	Int64      int64       = 1
	Uint       uint        = 1
	Uint8      uint8       = 1
	Uint16     uint16      = 1
	Uint32     uint32      = 1
	Uint64     uint64      = 1
	Uintptr    uintptr     = 1
	Float32    float32     = 1.1
	Float64    float64     = 1.1
	Complex64  complex64   = 1 + 2i
	Complex128 complex128  = 1 + 2i
	Array      [1]int      = [1]int{1}
	Chan       chan int    = make(chan int, 1)
	Func       func(int)   = func(a int) {}
	Interface  interface{} = 1
	Map        map[int]int = make(map[int]int)

	s                          = "a"
	Ptr         *string        = &s
	Slice       []int          = []int{1}
	String      string         = "a"
	Struct      SS             = SS{Name: "Tom", age: 20}
	UnsafePoint unsafe.Pointer = unsafe.Pointer(&s)
)

func Test1(t *testing.T) {
	fmt.Println(reflect.TypeOf(Bool))
	fmt.Println(reflect.TypeOf(Int))
	fmt.Println(reflect.TypeOf(Int8))
	fmt.Println(reflect.TypeOf(Int16))
	fmt.Println(reflect.TypeOf(Int32))
	fmt.Println(reflect.TypeOf(Int64))
	fmt.Println(reflect.TypeOf(Uint))
	fmt.Println(reflect.TypeOf(Uint8))
	fmt.Println(reflect.TypeOf(Uint16))
	fmt.Println(reflect.TypeOf(Uint32))
	fmt.Println(reflect.TypeOf(Uint64))
	fmt.Println(reflect.TypeOf(Uintptr))
	fmt.Println(reflect.TypeOf(Float32))
	fmt.Println(reflect.TypeOf(Float64))
	fmt.Println(reflect.TypeOf(Complex64))
	fmt.Println(reflect.TypeOf(Complex128))
	fmt.Println(reflect.TypeOf(Array))
	fmt.Println(reflect.TypeOf(Chan))
	fmt.Println(reflect.TypeOf(Func))
	fmt.Println(reflect.TypeOf(Map))
	fmt.Println(reflect.TypeOf(Interface))
	fmt.Println(reflect.TypeOf(Ptr))
	fmt.Println(reflect.TypeOf(Slice))
	fmt.Println(reflect.TypeOf(String))
	fmt.Println(reflect.TypeOf(Struct))
	fmt.Println(reflect.TypeOf(UnsafePoint))
}

type MyInt int

func Test2(t *testing.T) {
	a := MyInt(1)
	fmt.Println(reflect.TypeOf(a))
}

type MyInterface interface {
	Show()
}

type Show struct {
	Name string
	Age  int
}

func (s *Show) Show() {

}

func (s *Show) Hello() {

}

func Test3(t *testing.T) {
	b := &Show{}
	var c MyInterface
	c = b
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(c))
	fmt.Println(reflect.ValueOf(c).Type())

	fmt.Println(reflect.TypeOf(Show{}).Name())   // Show
	fmt.Println(reflect.TypeOf(Show{}).String()) // reflectt.Show
	fmt.Println(reflect.TypeOf([]int{1}).Size())
}

func Test4(t *testing.T) {
	a := &Show{}
	b := Show{}
	var c *Show
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)))
	fmt.Println(reflect.Indirect(reflect.ValueOf(b)))
	fmt.Println(reflect.Indirect(reflect.ValueOf(c)))
}

type MyInterface2 interface {
	Show()
	Hello()
}

func Test5(t *testing.T) {
	type A struct {
		B Show
	}
	a := A{B: Show{Name: "Tom", Age: 1}}
	fmt.Println(reflect.ValueOf(a).FieldByIndex([]int{0, 1}))

	b := &a
	fmt.Println(reflect.Indirect(reflect.ValueOf(b)).FieldByIndex([]int{0}))

	c := make(map[string]string)
	c["A"] = "abc"

	fmt.Println(reflect.ValueOf(c).MapIndex(reflect.ValueOf("A")))

	//	var aa MyInterface
	//	bb := &Show{Name:"Tom",Age:20}
	//	aa = bb
	//	dd := &aa
	//	ee := aa
	//
	//	fmt.Println(reflect.TypeOf(aa)) // *reflectt.Show
	//	fmt.Println(reflect.ValueOf(aa)) // &{Tom 20}
	//
	//	fmt.Println(reflect.TypeOf(bb).Name())
	//	fmt.Println(reflect.TypeOf(*bb).Name())
	//	fmt.Println(reflect.ValueOf(bb))
	//
	//	rv := reflect.ValueOf(dd).Elem()
	//	fmt.Println(rv, rv.Type(), rv.Type().Kind())
	//
	//	fmt.Println("-----------")
	//	fmt.Println(reflect.TypeOf(dd)) // *MyInterface
	//	fmt.Println(reflect.ValueOf(dd)) // ox
	//	fmt.Println(reflect.TypeOf(ee))  // *Show
	//	fmt.Println(reflect.ValueOf(ee)) // &{}
	//}
	//
	//func Test6(t *testing.T) {
	//	a := &Show{}
	//	fmt.Println(reflect.TypeOf(a)) // *main.Show
	//	fmt.Println(reflect.TypeOf(&a))// **main.Show
	//	var b MyInterface
	//	b = a
	//	fmt.Println(reflect.TypeOf(b)) // *main.Show
	//	fmt.Println(reflect.TypeOf(b).Elem().Kind()) // struct
	//	fmt.Println(reflect.TypeOf(&b)) // *main.MyInterface
	//	fmt.Println(reflect.TypeOf(&b).Elem()) // main.MyInterface
	//	fmt.Println(reflect.TypeOf(&b).Elem().Kind()) // interface
	//	var c MyInterface2
	//	c = a
	//	fmt.Println(reflect.TypeOf(c)) // *main.Show
	//	fmt.Println(reflect.TypeOf(c).Elem().Kind()) // struct
}

type I1 interface {
	F()
}

type I2 interface {
	I1
}

type I3 interface {
	I2
}

type I struct {
}

func (i *I) F() {

}

func Test7(t *testing.T) {
	i := &I{}
	var i1 I1 = i
	var i2 I2 = i1
	var i3 I3 = i2
	fmt.Println(reflect.TypeOf(i3))                // *main.I
	fmt.Println(reflect.TypeOf(&i3))               // *main.I3
	fmt.Println(reflect.TypeOf(&i3).Elem())        // main.I3
	fmt.Println(reflect.TypeOf(&i3).Elem().Kind()) // interface

	fmt.Println(reflect.ValueOf(i3))         // &{}
	fmt.Println(reflect.ValueOf(&i3))        // ox
	fmt.Println(reflect.ValueOf(&i3).Elem()) //
	fmt.Printf("%p", i)

}

func Test8(t *testing.T) {
	// 获取reflect.Value
	var a interface{}
	a = 1
	fmt.Println(getValue(a))

}

// 获取reflect.Value
func getValue(data interface{}) reflect.Value {
	if data == nil {
		return reflect.Value{}
	}
	v := reflect.ValueOf(data)
	if !v.IsValid() { // 没有类型和值信息
		return v
	}
	// 指针，接口
	if v.Kind() == reflect.Ptr && v.IsNil() {

	}
	return reflect.Value{}
}

func Test9(t *testing.T) {
	//var a = 2
	//v := reflect.ValueOf(&a).Elem()
	//if v.CanSet() {
	//	v.SetInt(3)
	//}
	//fmt.Println(a)
	//
	//var s = &Teacher{Name:"tom"}
	//tp := reflect.TypeOf(s).Elem()
	////val := reflect.ValueOf(s).Elem()
	//for i :=0;i < tp.NumField();i++ {
	//	field := tp.Field(i)
	//	//v := val.Field(i)
	//	fmt.Printf("%+v\n",field)
	//}
	//stp := reflect.TypeOf(s)
	//for i := 0; i < stp.NumMethod(); i++ {
	//	f := stp.Method(i)
	//	fmt.Printf("Name: %s, Type: %v, Func: %v, Index:%d, PkgPath: %s\n",f.Name,f.Type,f.Func,f.Index,f.PkgPath)
	//}
	//
	//val := reflect.ValueOf(s)
	//mf := val.MethodByName("Set")
	//mf.Call([]reflect.Value{reflect.ValueOf("jack")})
	//fmt.Println(s.Speak())
}

func Test10(t *testing.T) {
	var a MyInterface
	a = &Show{Name: "tom", Age: 18}
	fmt.Println(a)

	v := reflect.ValueOf(a)
	fmt.Println(v.CanAddr())

}

type i interface {
}

func Test11(t *testing.T) {
	var v i
	v = 2
	getValue := reflect.ValueOf(&v).Elem()
	fmt.Println(getValue.Kind())
}

func Test12(t *testing.T) {
	a := Show{
		Name: "Tom",
		Age:  18,
	}
	v := reflect.ValueOf(&a)
	v = reflect.Indirect(v)
	if v.CanAddr() {
		fmt.Println("CanAddr")
	}
	if v.CanSet() {
		fmt.Println("CanSet")
	}

	fmt.Println(a)
}

func Test13(t *testing.T) {
	a := make(chan int, 1)
	b := make(chan int, 1)
	fmt.Println(a == b)

	v := reflect.ValueOf(&a)
	v = reflect.Indirect(v)
	fmt.Println(v.Kind())
	fmt.Println(v.Type().Name())
	fmt.Println(v.CanAddr())
}

func indirect(v reflect.Value) reflect.Value {
	for {
		// Load value from interface, but only if the result will be
		// usefully addressable.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Ptr && !e.IsNil() {
				v = e
				continue
			}
		}
		if v.Kind() != reflect.Ptr {
			break
		}
		if v.IsNil() {
			return v
		}
		v = v.Elem()
	}

	return v
}

func indirect2(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return v
		}
		return indirect2(v.Elem())
	}
	return v
}

func indirectEnsure(v reflect.Value, newWhenNil bool) reflect.Value {
	if !v.IsValid() {
		return v
	}
	// If v is a named type and is addressable,
	// start with its address, so that if the type has pointer methods,
	// we find them.
	if v.Kind() != reflect.Ptr && v.Type().Name() != "" && v.CanAddr() {
		v = v.Addr()
	}

	for {
		if v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
			if v.IsNil() {
				if v.Kind() == reflect.Ptr {
					if newWhenNil {
						if v.CanSet() {
							v.Set(reflect.New(v.Type().Elem()))
						} else {
							return reflect.Value{}
						}
					} else {
						return reflect.Value{}
					}
				} else {
					return v
				}
			} else {
				v = v.Elem()
			}
		} else {
			break
		}
	}

	return v
}

func Test14(t *testing.T) {
	var a i
	v := reflect.ValueOf(&a)
	fmt.Println(v)
	fmt.Printf("%#v\n", indirect(v))
}
