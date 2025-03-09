package main

import "fmt"

func main() {
	// Basic for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For as while
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop
	// for {

	// }

	// For range
	numbers := []int{1, 2, 3, 4, 5}
	for index, num := range numbers {
		fmt.Println(index, num)
	}

	// Break the loop
	for i := 1; i <= 10; i++ {
		if i % 2 != 0 {
			continue
		}

		fmt.Println(i)
		
		if i > 5 {
			break
		}
	}
}