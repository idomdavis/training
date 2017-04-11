package main

import "fmt"

func foo() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("recovered: %v\n", x)
		}
	}()
	panic("kaboom")
}

func main() {
	foo()
	println("hello")
}
