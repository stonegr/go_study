package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "123ab='234'"
	r := strings.Index(a, "ab")
	quotation_mark_pos := strings.LastIndex(a, "'")
	fmt.Println(r, quotation_mark_pos)
	fmt.Println(a[r:quotation_mark_pos])
}
