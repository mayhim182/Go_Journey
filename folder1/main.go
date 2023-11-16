package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("Hello world")
	generalFunc()
	constant()
}

func generalFunc() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1:", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	page3()
}

func page3() {
	var a = "Initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

}

func constant() {
	fmt.Println(s)
	const n = 500000000

	const d = 3e20 / n

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	forLoopTut()
	arrayPrac()

}

func forLoopTut() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// Go arrays

func arrayPrac() {
	var a [5]int
	fmt.Println("Emp: ", a)
	a[4] = 100
	fmt.Println("Set: ", a)
	fmt.Println("Get: ", a[4])
	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	sum()
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

	nextInt := intSeq()

	fmt.Println(nextInt)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
