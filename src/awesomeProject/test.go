package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func loopTest() {
	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum)
	fmt.Println(sum1)
	fmt.Println(sum2)
}

//命名返回值
//Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
//直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
func namedReturn(sum int) (x, y int) {
	x = sum * 5 / 4
	y = sum - x
	return
}

//测试多值返回
func swapFake(a, b string) (string, string) {
	return b, a
}

//
func basicTypes() {
	//同导入语句一样，变量声明也可以“分组”成一个语法块。

	//bool
	//string
	//int  int8  int16  int32  int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//byte // uint8 的别名
	//rune // int32 的别名
	//    // 表示一个 Unicode 码点
	//float32 float64
	//complex64 complex128

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

//与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。
func typeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {

	swapA, swapB := "848", "shit"
	swapA, swapB = swapFake(swapA, swapB)
	fmt.Println("test multi-return ", swapA, swapB)
	fmt.Print("test namedReturn ")
	fmt.Println(namedReturn(134))
	fmt.Println("fuck u random ", rand.Intn(100))
	//不管是包中还是以后自己书写 大写变量都是外部可见 小写外部不可见
	fmt.Println(math.Pi)
	//numericConstants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//在 main 的 fmt.Println 调用开始前，两次对 pow 的调用均已执行并返回其各自的结果。
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
