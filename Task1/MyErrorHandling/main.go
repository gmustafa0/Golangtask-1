package main

import (
	"fmt"
	"strconv"
)

// 2. Implementing Error Interface
type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("cant divide %v by zero", de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend}
	}
	return dividend / divisor, nil
}

func main() {
	i, err := strconv.Atoi("69")

	if err != nil {
		fmt.Println("coudld'nt convert: ", err)
		return
	}
	fmt.Println("converted string is", i)

	//implemented error interface's driver code
	result, err := divide(10, 1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

}
