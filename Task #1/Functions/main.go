package main

import "fmt"

//simple function
func add(x, y int) int {
	return x + y
}

//function with multiple return types

func mulReturntypes(x, y, z int) (int, int, int) {
	return x, y, z
}

// Ignoring return values
func ignoreReturnval() (string, int) {
	return "Aman", 10
}

//implicit return function
func implicitreturn() (x int) {
	return
}

func main() {

	fmt.Println("Add function", add(1, 2))

	fmt.Println("This function returns 3 values:")
	fmt.Println(mulReturntypes(1, 2, 3))
	x, _ := ignoreReturnval()
	fmt.Println(x)
	implicitreturn()

}
