package main

import (
	"errors"
	"fmt"
)

var Var1 = 50
var isActive bool

func test() {

}

func main() {
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちは世界\n")

	// 定义变量
	// 隐式声明 var一般声明全局变量
	_, b := 34, "wangzijian"

	fmt.Println(b)
	fmt.Println(Var1)

	// 定义常量
	const PI = 3.1415926
	const prefix = "test"

	// 测试test
	test()

	// 定义复数
	var c complex64 = 5 + 5i
	fmt.Println("Value is : %v", c)

	// 字符串无法进行独立赋值
	s := "hello"
	// s[0] = 'c'是错误的
	c1 := []byte(s) // 将字符串改为byte类型
	c1[0] = 'c'
	// 在转换为字符串类型
	s2 := string(c1)
	fmt.Println(s2)

	// 字符串的拼接操作
	s3 := "hello"
	s4 := "world"
	a := s3 + s4
	fmt.Println(a)

	// 声明多行字符
	m := `hello
          world`
	fmt.Println(m)

	// 内置错误处理
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

}
