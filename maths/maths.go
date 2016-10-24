package maths

func Fibonacci(qty int) []int {
	var fib []int
	for i, f, s := 0, 0, 1; i < qty; i, f, s = i+1, s, f+s {
		fib = append(fib, f)
	}

	return fib
}

func Mean(numbers []float64) float64{
	sum := float64(0)
	for _, x := range numbers{
		sum += x
	}

	return sum/float64(len(numbers))
}
