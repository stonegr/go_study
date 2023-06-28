package my_moudle

import "fmt"

func Do_List() {
	// 数组
	// l0 := [5]string{"a", "0"}
	// fmt.Println(l0[1])
	// fmt.Println(l0)
	// cycle_arry(l0)

	// 切片
	// s0 := make([]string, 'a')
	s0 := make([]string, 2)
	// var s0 = []int{'a'}
	s0 = append(s0, "back")
	s0[0] = "0"
	fmt.Println(len(s0))
	fmt.Println(s0[0])
	fmt.Println(s0)

	s1 := make([]string, len(s0), cap(s0)*2)
	copy(s1, s0)
	fmt.Println(s1)

	s2 := []string{"1"}
	printSlice(s2)
	s2 = append(s2, "2")
	printSlice(s2)
	fmt.Println(s2)

	s3 := []int{1, 1, 1}
	newS := myAppend(s3)

	fmt.Println(s3)
	fmt.Println(newS)

	s3 = newS

	myAppendPtr(&s3)
	fmt.Println(s3)

	myAppend_do(s3)
	fmt.Println(s3)
}

// func cycle_arry(l [5]string) {
// 	for k, v := range l {
// 		fmt.Printf("%d: %s\n", k, v)
// 	}
// }

func printSlice(x []string) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppend_do(s []int) {
	// 可以改变 slice 底层数组元素值。
	s[0] = 0
	// return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
}
