package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := []int{10, 20, 30}
	fmt.Println("a: ", a)
	a = push(a, 100)
	fmt.Println("a: ", a)
	a = pop(a)
	fmt.Println("a: ", a)
	a = unshift(a, 1000)
	fmt.Println("a: ", a)
	a = shift(a)
	fmt.Println("a: ", a)
	a = insert(a, 1000, 2)
	fmt.Println("a: ", a)
	a = remove(a, 2)
	fmt.Println("a: ", a)
}

func push(a []int, n ...int) (s []int) {
	s = append(a, n...)
	return
}

func pop(a []int) []int {
	return a[:len(a)-1]
}

func unshift(a []int, n int) []int {
	return append([]int{n}, a...)
}

func shift(a []int) []int {
	return a[1:]
}

func insert(a []int, n, i int) []int {
	a = append(a, 0)
	a = append(a[:i+1], a[i:len(a)-1]...)
	a[i] = n
	return a
}

func remove(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
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
