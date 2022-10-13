https://golang.google.cn/tour/basics/15

Constants cannot be declared using the := syntax.

函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。

没有明确初始值的变量声明会被赋予它们的 零值。
数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）。

类型转换
表达式 T(v) 将值 v 转换为类型 T。
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

常量的声明与变量类似，只不过是使用 const 关键字。
常量不能用 := 语法声明。

Go 只有一种循环结构：for 循环。
sum := 1
for ; sum < 1000; {
	sum += sum
}
#for 是 Go 中的 “while”
sum := 1
for sum < 1000 {
	sum += sum
}

Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

defer 语句会将函数推迟到外层函数返回之后执行。
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}


Go 拥有指针。指针保存了值的内存地址。
var p *int

i := 42
p = &i
fmt.Println(*p) // 通过指针 p 读取 i

与 C 不同，Go 没有指针运算。


一个结构体（struct）就是一组字段（field）。
结构体字段使用点号来访问。
结构体字段可以通过结构体指针来访问。
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
v2 = Vertex{X: 1}  // Y:0 被隐式地赋予

数组
类型 [n]T 表示拥有 n 个 T 类型的值的数组。
var a [10]int ，因此数组不能改变大小。
primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	切片
每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。
类型 []T 表示一个元素类型为 T 的切片。
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

切片并不存储任何数据，它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

与它共享底层数组的切片都会观测到这些修改。


切片文法
切片文法类似于没有长度的数组文法。
s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

在进行切片时，你可以利用它的默认行为来忽略上下界。
对于数组

var a [10]int
来说，以下切片是等价的：

a[0:10]
a[:10]
a[0:]
a[:]


nil 切片
切片的零值是 nil。


切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。
a := make([]int, 5)  // len(a)=5

要指定它的容量，需向 make 传入第三个参数：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

切片可包含任何类型，甚至包括其它的切片。


向切片追加元素
为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。内建函数的文档对此函数有详细的介绍。

func append(s []T, vs ...T) []T

Range
for 循环的 range 形式可遍历切片或映射。

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

range（续）
可以将下标或值赋予 _ 来忽略它。
for i, _ := range pow
for _, value := range pow


