package main

import (
	"fmt"
	"os"
)

func GetPathfiles(path string) []string {
	var all_file_slice []string
	files, _ := os.ReadDir("./cs")
	fmt.Println(files[0].Name())
	// sort.SliceStable(files, func(i, j int) bool {
	// 	return files[i].ModTime().Before(files[j].ModTime())
	// })
	// for _, f := range files {
	// 	all_file_slice = append(all_file_slice, f.Name())
	// }
	return all_file_slice
}

func main() {
	r := GetPathfiles("./cs")
	// for _, f := range r {
	// 	fmt.Println(f)
	// }
	fmt.Println(r[0])
}
