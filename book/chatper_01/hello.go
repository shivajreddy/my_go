package main

import "fmt"

func main() {

	fmt.Println("Hello %s there", "world")
	var x int
	x = 21
	fmt.Println("x=", x)
	x += 1
	fmt.Println("x=", x)

	var f bool
	f = false
	fmt.Println(f)
	f = true
	fmt.Println(f)

	a := 10
	fmt.Println("a=", a)
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("Hello World")
}
