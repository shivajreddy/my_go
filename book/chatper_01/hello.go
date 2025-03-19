/*
Create a Linked List in GO
*/
package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

// ENTRY POINT
func main() {
	n := Node{}

	fmt.Println(n.val)
	fmt.Println(n)
	fmt.Println(nil)
	fmt.Println(n.next)
}
