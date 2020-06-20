package learn

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	a := []int{1, 2, 3}
	b := a
	b = append(b[:1], b[1+1:]...)
	fmt.Println(a)
	fmt.Println(b)
}
