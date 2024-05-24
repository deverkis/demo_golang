## 常量
常量的声明与变量类似，只不过使用 const 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 := 语法声明。

## 数值常量
数值常量是高精度的 值。

一个未指定类型的常量由上下文来决定其类型。

```bash
package main

import "fmt"

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 16
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

## for 循环
Go 只有一种循环结构：for 循环。

Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的。

### 无限循环
如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
for {
	}

## if 判断
Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( )，而大括号 { } 则是必须的。

和 for 一样，if 语句可以在条件表达式前执行一个简短语句。

```bash
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```

## switch 分支
switch 语句是编写一连串 if - else 语句的简便方法。它运行第一个 case 值 值等于条件表达式的子句。

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只会运行选定的 case，而非之后所有的 case。 

在效果上，Go 的做法相当于这些语言中为每个 case 后面自动添加了所需的 break 语句。在 Go 中，除非以 fallthrough 语句结束，否则分支会自动终止。 

Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不限于整数。

## 无条件 switch
无条件的 switch 同 switch true 一样。

这种形式能将一长串 if-then-else 写得更加清晰。

```bash
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}
```

## defer 推迟
defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

## defer 栈
推迟调用的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的调用会按照后进先出的顺序调用。

## 指针
Go 拥有指针。指针保存了值的内存地址。

类型 *T 是指向 T 类型值的指针，其零值为 nil。