package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	in := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	
	type info struct{ count, length int }
	calculated := map[string]info{}
	for _, s := range in {
		calculated[s] = info{
			count:  strings.Count(s, "a"),
			length: utf8.RuneCountInString(s),
		}
	}
	
	sort.Slice(in, func(i, j int) bool {
		inf1, inf2 := calculated[in[i]], calculated[in[j]]
		if inf1.count != inf2.count {
			return inf1.count > inf2.count
		}
		return inf1.length > inf2.length
	})

	fmt.Println(in) 
}
