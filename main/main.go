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

}
