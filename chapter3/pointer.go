package main

import "fmt"

func main() {
	n := 123
	p := &n
	fmt.Println("n: ", n)
	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)

}
