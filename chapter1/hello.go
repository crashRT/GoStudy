package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	x := 5
	switch x {
	case f(1):
		fmt.Println("1")
	case f(2):
		fmt.Println("2")
	case f(3):
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}

func f(n int) int {
	fmt.Println("No. ", n)
	return n
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
