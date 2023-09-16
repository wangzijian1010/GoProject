package main

// 如果"fmt" 前面加个. 等价于c++的using namespace
//  也可以取别名 f "fmt"
//

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

var Var1 = 50
var isActive bool

func test() {

}

// 定义函数
func test1(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 定义函数返回多个值类似python
// 返回多个值的时候需要在括号多加一个int
func test2(a int, b int) (int, int) {
	return a + b, a * b
}

// 变参 不同的数量参数
// 不同数量的全为int
func myfunc(arg ...int) {

}

// 传指针 go默认函数为值copy
func add1(a *int) int {
	*a = *a + 1
	return *a
}

type testInt func(int) bool // 返回一个Bool

// 定义一个示例操作函数，将给定的整数加倍
func double(number int) int {
	return number * 2
}

// 定义另一个示例操作函数，将给定的整数加一
func increment(number int) int {
	return number + 1
}

// 定义一个接受函数作为参数的函数
func applyOperation(numbers []int, operation func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = operation(num)
	}
	return result
}

// 传入结构体
func Older(p1 person, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

// 结构体内继续定义结构体
type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，那么默认 Student 就包含了 Human 的所有字段
	speciality string
}

// method的实现
// method的实例
// 类似C++的构造函数

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

//
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // a slice of boxes

func (b *Box) SetColor(c Color) {
	b.color = c
}
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

type Human1 struct {
	name  string
	age   int
	phone string
}

type Student1 struct {
	// 继承human结构体
	Human1
	school string
	loan   float32
}

type Employee struct {
	Human1
	company string
	money   float32
}

// Human实现sayhi方法
func (h *Human1) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human1) Sing() {
	fmt.Println("La la, la la la, la la la la la...")
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
		fmt.Println(err)
	}

	// 数组定义
	var arr [10]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println("the second num is ", arr[1])

	// go语言传入数组的均为传值而不是指针
	arr2 := [3]int{1, 2, 4}
	fmt.Println("the second num is ", arr2[1])

	// 省略长度也可以赋值
	arr3 := [...]int{3, 5, 6}
	fmt.Println("the second num is ", arr3[1])

	// 声明二维数组
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Println("the second num is ", doubleArray[1][2])

	// slice类似动态数组类似cpp的vector
	// 使用append增加元素
	slice := []string{"a", "b", "c"}
	fmt.Println(slice)

	// map为键值对的格式
	// 声明一个 key 是字符串，值为 int 的字典,这种方式的声明需要在使用之前使用 make 初始化
	// var numbers0 map[string]int

	// 另外一种声明方式
	// key为字符串 int为值
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	fmt.Println(numbers["two"])

	// 初始化赋值
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	fmt.Println(rating["C"])

	csharpRating, ok := rating["C++"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("doesn't exist ")

	}

	// 删除键值为C的值
	delete(rating, "C")

	maxInput := test1(2, 3)
	fmt.Println(maxInput)

	sum1, mul1 := test2(3, 7)
	fmt.Println(sum1, mul1)

	// 传入指针地址
	testInput := 7
	add1(&testInput)
	fmt.Println(testInput)

	// defer采用后进先出 类似栈
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	// 函数可以作为类型传入做参数
	// 要处理的整数切片
	numbers4 := []int{1, 2, 3, 4, 5}
	doubledNumbers := applyOperation(numbers4, double)
	incrementedNumbers := applyOperation(numbers4, increment)
	fmt.Println("Original Numbers:", numbers)
	fmt.Println("Doubled Numbers:", doubledNumbers)
	fmt.Println("Incremented Numbers:", incrementedNumbers)

	// 使用结构体
	var P person // P 是person类型的变量 类似C++的实例化
	P.age = 1
	P.name = "wangzijian"

	// 顺序赋值
	var P1 person
	P1 = person{"Tom", 25}
	fmt.Println(P1)

	// 结构体传入
	tb_Older, tb_diff := Older(P, P1)
	fmt.Println(tb_Older)
	fmt.Println(tb_diff)

	// 初始化结构体中的结构体
	mark := Student{Human{"wangzijian", 25, 120}, "computer science"}
	fmt.Println("His name is ", mark.name)

	r1 := Rectangle{12, 2}
	r2 := Rectangle{13, 3}
	c2 := Circle{22}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c2 is: ", c2.area())

}
