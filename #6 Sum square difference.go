package main

import (
	"fmt"
	"math"
)
//comments are for debugging

func main() {
	var sum float64
	for i := 1; i < 101; i++ {
		sum += float64(i)
		//fmt.Println("plus ",i,"  ",sum)
	}
	sum = math.Pow(sum, 2)
	//fmt.Println("pow of sum", sum)
	for i := 1; i < 101; i++ {
		sum -= math.Pow(float64(i), 2)
		// fmt.Println("minus ",i,"   ",sum)
	}
	fmt.Println(sum)
}
