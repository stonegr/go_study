package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.Split("wewe/w/w/ssds/sdsdssddvf", "/")
	xx := x[len(x)-1]
	xxx := strings.Join(x[:len(x)-1], "/")
	fmt.Println(xx, xxx)
}
