package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("One argument is needed")
		os.Exit(0)
	}
	args := os.Args[1:]
	switch args[0] {
	case "type":
		typeDemo()
	case "select":
		selectDemo()
	}
}
func typeDemo() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T \n", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}

func selectDemo() {
//	var c2, c3 chan int
	var i1, i2 int
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)
	go func() { c1 <- 1 }()
	go func() { c3 <- 3 }()

	for{
		time.Sleep(2 * time.Second)
		select {
		case i1 = <-c1:
			fmt.Printf("received %d from c1\n", i1)
		case c2 <- i2:
			fmt.Printf("sent %d to c2\n", i2)
		case i3, ok := (<-c3):  // same as: i3, ok := <-c3
			if ok {
				fmt.Printf("received %d from c3\n", i3)
			} else {
				fmt.Printf("c3 is closed\n")
			}
		default:
			fmt.Printf("no communication\n")
		}
	}

}
