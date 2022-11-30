package main

import (
	"fmt"
)

// recursive function

func print_num(n int) {
	

	// Calling the function with decremented value
	if n > 1 {
		
		// the function
		print_num(n-7)
		print_num(n-5)
		
		// printing is done at
		// returning time
		fmt.Println(n)
	}
}

// main function
func main() {
	
//    Prints numbers on question
	print_num(9)
}
