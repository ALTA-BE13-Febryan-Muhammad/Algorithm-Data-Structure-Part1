package main

import "fmt"

func PrimeX(number int) int {
	var x = []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	// fmt.Println(x[number])
	return int(x[number])

}

func main() {
	fmt.Println(PrimeX(1))  //2
	fmt.Println(PrimeX(5))  //11
	fmt.Println(PrimeX(8))  //19
	fmt.Println(PrimeX(9))  //23
	fmt.Println(PrimeX(10)) //29
}
