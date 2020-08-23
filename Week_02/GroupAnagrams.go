package main

import (
	"fmt"
	"sort"
)

func SortString(s string) string {
	r := []rune(s)

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 0 {
		return [][]string{}
	}

	r := map[string][]string{}

	for _, str := range strs {
		tmp := SortString(str)

		if list, ok := r[tmp]; ok {
			r[tmp] = append(list, str)
		} else {
			r[tmp] = []string{str}
		}
	}

	var a [][]string
	for _, strings := range r {
		a = append(a, strings)
	}

	return a
}

func main() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	l := groupAnagrams(a)

	for _, v := range l {
		fmt.Println(v)
	}
}
