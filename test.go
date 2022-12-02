package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {

	// words := []string{"apple", "pie", "apple", "red", "red", "red"}

	// Collecting the strings in the arraylist and their count in a histogram.

	for val := 1; val < 6; val++ {
		fmt.Printf("%d", val)
		switch {
		case val == 1:
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
			fmt.Printf("%s", strhi)
			fmt.Printf("\n")

		case val == 2:
			words := []string{"keyboard", "go", "js", "cpp", "cpp", "java"}

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
			fmt.Printf("%s", strhi)
			fmt.Printf("\n")

			break
		case val == 3:
			words := []string{"country", "town", "city", "province", "state", "province"}

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
			fmt.Printf("%s", strhi)
			fmt.Printf("\n")

			break
		case val == 4:
			words := []string{"football", "rugby", "basketball", "rugby", "tennis", "swimming"}

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
			fmt.Printf("%s", strhi)
			fmt.Printf("\n")

			break
		case val == 5:
			words := []string{"music", "art", "entertainment", "acting", "art", "art"}

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
			fmt.Printf("%s", strhi)
			fmt.Printf("\n")

			break
		default:
			t.Errorf("String returned is incorrect!")
		}
	}

}
