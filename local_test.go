package learn

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	a := []int{1, 2, 3}
	b := a
	b = append(b[:1], b[1+1:]...)
	fmt.Println(a)
	fmt.Println(b)
}

func Test2(t *testing.T) {
	log.Print("1")
	log.Printf("%d", 2)
	log.Println("3")
	log.Fatal("4")
}

func Test3(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	fmt.Println(len(ch), cap(ch))
}

func Test4(t *testing.T) {
	ln, err := net.Listen("tcp", "192.168.2.50:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	ln.Close()
	ln.Close()
	fmt.Println("closed")
}

func Test5(t *testing.T) {
	ch := make(chan int, 1)
	close(ch)
	select {
	case ch <- 1:
	default:
		fmt.Println("hh")
	}
}

func Test6(t *testing.T) {
	f := func() {
		go func() {
			defer func() {
				recover()
			}()
			panic("test")
		}()
	}

	f()
	time.Sleep(10 * time.Second)
	fmt.Println("end")
}

func Test7(t *testing.T) {
	ln, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("accept error", err)
				return
			}
			fmt.Println("new client", conn.RemoteAddr())
			go func() {
				for {
					var data []byte
					if _, err := conn.Read(data); err != nil {
						fmt.Println("conn read error", err)
						return
					}
					fmt.Println(string(data))
				}
			}()
		}
	}()
	go func() {
		time.Sleep(time.Second)
		ln.Close()
		fmt.Println("close conn")
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("close")
}

func Test8(t *testing.T) {
	ch := make(chan struct{}, 1)
	//ch <- struct{}{}
	close(ch)
	_, ok := <-ch
	fmt.Println(ok)
}
