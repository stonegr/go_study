package main

import (
	"fmt"
	"strings"
)

func main() {
	// a := "https://cdn2.efhie.com/fvod/13cd72237aecf77a4900ed8306b1a339e1f57ca6f837c3f2c2e354a43160c47ef30c5cbce66f15450e77429e0ba7b8e9576d797e0bb592347f8d5b4665aa5d51578f567936e2b754fd0d1d19b0a07e67ce6f1263a1d2fbd0.ts"
	a := "13cd72237aecf77a4900ed8306b1a339e1f57ca6f837c3f2c2e354a43160c47ef30c5cbce66f15450e77429e0ba7b8e9576d797e0bb592347f8d5b4665aa5d51578f567936e2b754fd0d1d19b0a07e67ce6f1263a1d2fbd0.ts"
	aa := strings.SplitAfter(a, "/")
	fmt.Println(len(aa), aa)
	fmt.Println(strings.Join(aa[:len(aa)-1], ""))
	b := "123"
	fmt.Println(strings.ReplaceAll(b, "", "1"))
}
