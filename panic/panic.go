package main

func foo() {
	panic("kaboom")
}

func main() {
	foo()
	println("hello")
}
