package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dup3()
}

func combine(reader io.Reader, m map[string]int) {
	input := bufio.NewScanner(reader)
	for input.Scan() {
		m[input.Text()]++
	}
}

func printRes(m map[string]int) {
	for k, v := range m {
		fmt.Println(v, k)
	}
}

func dup1() {
	m := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		if s == "end" {
			break
		}
		m[s]++
	}
	printRes(m)
}

func dup2() {
	if len(os.Args) < 2 {
		fmt.Println("no file name")
		return
	}
	m := make(map[string]int)
	for _, v := range os.Args[1:] {
		file, err := os.Open(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		combine(file, m)
		file.Close()
	}
	printRes(m)
}

func dup3() {
	if len(os.Args) < 2 {
		fmt.Println("no file name")
		return
	}
	m := make(map[string]int)
	for _, v := range os.Args[1:] {
		by, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		for _, vv := range strings.Split(string(by), "\n") {
			m[vv]++
			if m[vv] > 1 {
				fmt.Println("-->", v)
			}
		}
	}
	printRes(m)
}
