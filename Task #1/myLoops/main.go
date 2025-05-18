package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}

	print("\n\n")
	i := 0
	for i < 5 {
		fmt.Println(i + 1)
		i++
	}

	numbers := []int{5, 8, 14}
	for index, value := range numbers {
		fmt.Printf("slice: index=%d, value=%d\n", index, value)
	}
}
