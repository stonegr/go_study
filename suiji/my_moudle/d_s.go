package my_moudle

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func Do_d() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "ab"})

	d0 := Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407}

	// 访问结构体成员
	fmt.Println(d0.book_id)

	// 结构体作为参数
	printBook(d0)
	fmt.Println(d0)

	// 传递指针
	printBookPointer(&d0)
	fmt.Println(d0)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	book.book_id = 3
	fmt.Printf("Book book_id : %d\n", book.book_id)
	book.title = "title"
}

func printBookPointer(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
	book.title = "title"
}
