package main

import "fmt"

func main() {
	studentsAge := map[string]int{
		"solomon": 19,
		"john":    20,
		"janet":   15,
		"daniel":  16,
		"mary":    18,
	  }

	  for name, age := range studentsAge {
		fmt.Println(name, age)
	  }
}