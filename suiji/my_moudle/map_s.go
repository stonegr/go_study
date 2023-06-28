package my_moudle

import "fmt"

// Map 是一种无序的键值对的集合。指针类型，会改变值
func Do_map() {
	// make创建,initialCapacity初始容量
	// m0 := make(map[string]int32)
	// 使用字面量创建 Map
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 300000000000000,
	}
	fmt.Println(m1)
	printMap(m1)

	// 读取值
	fmt.Println(m1["apple"])
	// 设置值
	m1["banana"] = 7
	printMap(m1)
	// 增加值
	m1["append"] = 8
	printMap(m1)
	// 删除值
	delete(m1, "append")
	printMap(m1)

}

func printMap(d map[string]int) {
	fmt.Println("-------------------")
	for k, v := range d {
		fmt.Println(k, "-", v)
	}
}
