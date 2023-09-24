package main

import (
	"fmt"
	"net/url"
)

func main() {
	// a := "123ab='234'"
	// r := strings.Index(a, "ab")
	// quotation_mark_pos := strings.LastIndex(a, "'")
	// fmt.Println(r, quotation_mark_pos)
	// fmt.Println(a[r:quotation_mark_pos])

	a := "https://baidu.com/aa"
	// r, _ := url.JoinPath(a, "wer")
	// r, _ := url.JoinPath(a, "/wer")

	r, _ := url.JoinPath(a, "https://baidu.com")
	fmt.Println(r)
}
