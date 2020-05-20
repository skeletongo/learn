package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Address struct {
	City string
	Area string
}

type Student struct {
	Address
	Name string
	Age  int
}

func (this Student) Say() {
	fmt.Println("hello, i am ", this.Name, "and i am ", this.Age)
}

func (this Student) Hello(word string) {
	fmt.Println("hello", word, ". i am ", this.Name)
}

func main() {
	var student = new(Student)
	student.Name = "Tom"
	student.Age = 18
	fmt.Println(SturctInfo(*student))
}

// 获取结构体信息
func SturctInfo(obj interface{}) string {
	buf := strings.Builder{}
	// 类型
	t := reflect.TypeOf(obj)
	buf.WriteString(fmt.Sprintln("type:",t.Name()))
	if t.Kind() != reflect.Struct {
		fmt.Println("this obj is not a struct, it is", t.Kind())
		return ""
	}
	v := reflect.ValueOf(obj)
	// 字段信息
	for i := 0; i < t.NumField(); i++ {
		tt := t.Field(i)
		vv := v.Field(i).Interface()
		buf.WriteString(fmt.Sprintln(tt.Name, tt.Type, vv))
		t1 := reflect.TypeOf(vv)
		if t1.Kind() == reflect.Struct {
			buf.WriteString(SturctInfo(vv))
		}
	}
	// 方法信息
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		buf.WriteString(fmt.Sprintln(m.Name,m.Type))
	}
	return buf.String()
}