package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

var c, python, java = "c", "python", "java"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return x, y
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println("hellow", time.Now())
	fmt.Println(math.Pi)
	fmt.Println(add(123, 222))
	a, b := swap("sdf", "abc")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i = true
	fmt.Println("sdfsdf", i, python, c, java)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	cc := 0.2345
	fmt.Println(int(cc))

	fmt.Println(Small, needInt(Small))
	fmt.Println(Small, needFloat(Small))
	fmt.Println(needFloat(Big))
}
