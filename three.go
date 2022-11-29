package main

import (
	"fmt"
)

func main() {
	words := []string{"apple", "pie", "apple", "red", "red", "red"}

	
	histo := make(map[string]int)
	for _, str := range words {
		histo[str]++
	}

	
	valhi := 0
	strhi := ""
	for k, v := range histo {
		if v > valhi {
			valhi = v
			strhi = k
		}
	}
	fmt.Printf( "%s" , strhi)
}
