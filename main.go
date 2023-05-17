package main

import "fmt"

func main() {
	l1 := []int8{1, 2}
	fmt.Println(l1)
	fmt.Println(&l1)

	l1 = append(l1, 3)
	fmt.Println(l1)
}
