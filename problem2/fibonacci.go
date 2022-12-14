package main

import "fmt"

func Fibonacci(number int) int {
	// var n int
	if number <= 1 {
		return number
	}
	return Fibonacci(number-1) + Fibonacci(number-2)
	// return Fibonacci(number)
}

// func Fibonacci(n int) int {
// 	f := make([]int, n+1, n+2)
// 	if n < 2 {
// 		f = f[0:2]
// 	}
// 	f[0] = 0
// 	f[1] = 1
// 	for i := 2; i <= n; i++ {
// 		f[i] = f[i-1] + f[i-2]
// 	}
// 	return f[n]
// }

func main() {

	fmt.Println(Fibonacci(0))  // 0
	fmt.Println(Fibonacci(2))  // 1
	fmt.Println(Fibonacci(9))  // 34
	fmt.Println(Fibonacci(10)) // 55
	fmt.Println(Fibonacci(12)) // 144
	fmt.Println(Fibonacci(20)) //

}

// f(9)=f(8+7) // 8? 7?(6+5)
// f(8)=f(7+6) // 7? 6?(5+4)
// f(7)=f(6+5) // 6? 5?(4+3)
