package main

import (
	"fmt"
)

func print_num(n int) {

	if n > 1 {
		

		print_num(n-7)
		print_num(n-5)

		fmt.Println(n)
	}
}


func main() {

	print_num(9)
}
