package main

import (
	"strings"
	"testing"
)

func add(arr []string) string {
	s := ""
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

func join(arr []string) string {
	return strings.Join(arr, "")
}

func buf(arr []string) string {
	buf := strings.Builder{}
	for i := 0; i < len(arr); i++ {
		buf.WriteString(arr[i])
	}
	return buf.String()
}

var arr = []string{"a", "b", "c"}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(arr)
	}
	// BenchmarkAdd-20    	10000000	       139 ns/op
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join(arr)
	}
	// BenchmarkJoin-20    	20000000	        69.3 ns/op
}

func BenchmarkBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf(arr)
	}
	// BenchmarkBuf-20    	20000000	        63.0 ns/op
}
