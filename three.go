package main

import (
	"fmt"
)

func main() {
	words := []string{"apple", "pie", "apple", "red", "red", "red"}

	// Collecting the strings in the arraylist and their count in a histogram.
	histo := make(map[string]int)
	for _, str := range words {
		histo[str]++
	}

	// Scanning the histogram to find the string with the highest count.
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
