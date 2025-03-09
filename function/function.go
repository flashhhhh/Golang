package main

import (
	"errors"
	"fmt"
)

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Multiple returns function
func addSubtract(a, b int) (int, int) {
	return a + b, a - b
}

// Error handling
func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil
}

// Exception handling
func canBeDivided(a, b float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()

	if b == 0 {
		panic("Division by zero")
	}

	fmt.Println("Can be divided")
}

func main() {
	// Variadic function
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20, 30, 40))
	fmt.Println(sum())

	// Multiple returns function
	add, sub := addSubtract(10, 5)
	fmt.Println(add, sub)

	// Error handling
	result, err := division(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Exception handling
	canBeDivided(10, 0)
	canBeDivided(10, 2)
}