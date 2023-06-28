package my_moudle

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func Do_sting() {
	// 字符相加
	a := "aa"
	b := "bb"
	a = a + b
	fmt.Println(a)

	// 数组转字符串
	// 切片
	var s []string
	for i := 0; i < 9; i++ {
		s = append(s, strconv.Itoa(i))
	}
	fmt.Println(strings.Join(s, ""))

	// 比上面效率更高一点
	var buffer bytes.Buffer
	for i := 0; i < 9; i++ {
		buffer.WriteString(strconv.Itoa(i))
	}
	fmt.Println(buffer.String())

	// 字符串截取 不含中文
	sa := "abcdefg"
	// 左闭右开
	sa = string([]byte(sa)[1:])
	fmt.Println(sa)

	// 字符串截取 含中文
	sa2 := "a你好bcdefg"
	// 左闭右开,rune代表int32
	sa2 = string([]rune(sa2)[1:3])
	fmt.Println(sa2)

	// 字符串替换
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", 2))  // aBaACEDF
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", -1)) // aBaaCEDF

	// 字符串转大小写
	fmt.Println(strings.ToUpper("abaacedf")) //ABAACEDF
	fmt.Println(strings.ToLower("abaacedf")) //ABAACEDF

	// 字符串包含 区分大小写
	fmt.Println(strings.Contains("hello world", "lo"))

	// 字符串包含 子串str中的任何一个字符 区分大小写
	fmt.Println(strings.ContainsAny("hello world", "w"))

	info := " 1 2 3 "
	res := strings.TrimSpace(info)
	//把字符串以空格分割成字符串数组
	str_arr := strings.Split(res, " ")
	//计算字符串数组的长度
	count := len(str_arr)
	fmt.Println(count)

	// 类型转换
	// 字符串转整数
	s2 := "123"
	num, _ := strconv.Atoi(s2)
	fmt.Println(num)
	// 整数转字符串
	i0 := 123
	s3 := strconv.Itoa(i0)
	fmt.Println(s3)

}
