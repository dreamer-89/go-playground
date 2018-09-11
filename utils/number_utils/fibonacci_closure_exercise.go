package number_utils

import "fmt"

func fibonacci() func() int {
	i, a, b := 0, 0, 1
	return func() int {
		i = a
		a = a + b
		b = i
		return i
	}
}

// PrintFibonacci prints first 10 Fibonacci numbers.
func PrintFibonacci() {
	f := fibonacci()
	fmt.Print("first 10 of Fibonacci numbers: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	fmt.Print("\n")
}