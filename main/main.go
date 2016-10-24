package main

import (
	"../maths"
	"fmt"
)

func main() {
	fibs := maths.Fibonacci(20)

	for _,val := range fibs {
		fmt.Println(val)
	}


	fmt.Println("Finding average for [1, 2, 3, 4]")

	numbers := []float64{1, 2, 3, 4}
	avg := maths.Average(numbers)

	fmt.Println(avg)
}
