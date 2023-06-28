package my_moudle

import (
	"errors"
	"fmt"
)

// 接口的返回类型不是强制指定的
type Phone interface {
	call()
	cs() int
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (f NokiaPhone) cs() string {
	fmt.Println("I am Nokia cs")
	return "I am Nokia cs"
}

// 有参数的接口,这里的类型不强制，有没有这个都没有关系
type Shape interface {
	area() string
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	} else {
		return 1, nil
	}
	// 实现
}

type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func Do_jk() {
	// 无参数
	// p0 := new(NokiaPhone)
	p0 := NokiaPhone{}
	p0.call()
	p0_cs := p0.cs()
	fmt.Println(p0_cs)

	// 有参数
	p1 := Rectangle{1, 2}
	fmt.Println(p1.area())

	// 错误处理
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// 接口错误处理
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}
