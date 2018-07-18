package main

import (
	"fmt"
	"math"
	"runtime"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

type user struct {
	name int
	age  int
}

func main() {
	sum := 0
	i := 0
	for i < 100 {
		sum++
		i++
	}
	fmt.Println(sum)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	p := []int{2, 3, 5, 7, 11, 13}

	p = append(p, 12)
	print(p)
	fmt.Println("counting")
	for i, v := range p {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	fmt.Println(user{1, 2})
}
